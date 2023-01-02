package agent

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
)

type ProofStatus = agency.ProtocolStatus_PresentProofStatus
type ProofAttribute = agency.Protocol_Proof_Attribute
type ProofPredicate = agency.Protocol_Predicates_Predicate
type ProofPresentation = agency.Question_ProofVerifyMsg

type PresentProofState int

const (
	StateProofProposal     PresentProofState = 1
	StateProofRequest      PresentProofState = 2
	StateProofPresentation PresentProofState = 3
	StateProofDone         PresentProofState = 4
)

func (e PresentProofState) String() string {
	switch e {
	case StateProofProposal:
		return "PROOF_PROPOSAL"
	case StateProofRequest:
		return "PROOF_REQUEST"
	case StateProofPresentation:
		return "PROOF_PRESENTATION"
	case StateProofDone:
		return "PROOF_DONE"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type proofData struct {
	id          string
	questionID  string
	clientID    string
	actualState PresentProofState
	verifier    bool
}

type Proof struct {
	ID         string `json:"referent"`
	ProofDefID string `json:"proof_def_id"`
	SchemaID   string `json:"schema_id"`
}

type ProofStore struct {
	agent *AgencyClient
	store map[string]proofData
	mtx   sync.RWMutex
}

func InitProofs(a *AgencyClient) *ProofStore {
	return &ProofStore{
		agent: a,
		store: make(map[string]proofData),
	}
}

func (s *ProofStore) HandleProofNotification(notification *agency.Notification) (err error) {
	defer err2.Handle(&err)

	// Proof success
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}

			status := try.To1(s.agent.ProtocolClient.Status(context.TODO(), protocolID))

			if status.State.State == agency.ProtocolState_OK {
				proof := status.GetPresentProof()
				log.Printf("Proof ready %v\n", proof)

				// TODO: role in notification should indicate if we are holder or not
				verifier, _ := try.To2(s.GetProof(notification.ProtocolID))

				data := &proofData{
					id:          protocolID.ID,
					verifier:    verifier,
					actualState: StateProofDone,
				}
				try.To(s.addProofData(protocolID.ID, data))
			}
		}

		// Proof request received
	} else if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED &&
		notification.GetRole() == agency.Protocol_ADDRESSEE {
		data := &proofData{
			id:          notification.ProtocolID,
			verifier:    false,
			actualState: StateProofRequest,
		}
		try.To(s.addProofData(notification.ProtocolID, data))
	}
	return nil
}

func (s *ProofStore) HandleProofQuestion(question *agency.Question) (err error) {
	defer err2.Handle(&err)

	if question.TypeID == agency.Question_PROOF_VERIFY_WAITS {
		data := &proofData{
			id:          question.Status.Notification.ProtocolID,
			questionID:  question.Status.Notification.ID,
			clientID:    question.Status.ClientID.ID,
			verifier:    true,
			actualState: StateProofPresentation,
		}

		proof := question.GetProofVerify()
		log.Printf("Proof presented %v\n", proof)

		try.To(s.addProofData(data.id, data))
	} else if question.TypeID == agency.Question_PROOF_PROPOSE_WAITS {
		data := &proofData{
			id:          question.Status.Notification.ProtocolID,
			questionID:  question.Status.Notification.ID,
			clientID:    question.Status.ClientID.ID,
			verifier:    true,
			actualState: StateProofProposal,
		}

		proof := question.GetProofVerify()
		log.Printf("Proof proposal %v\n", proof)

		try.To(s.addProofData(data.id, data))

		// just accept proof propose directly
		log.Printf("Accept proof proposal with the thread id %s, question id %s", question.Status.ClientID.ID, question.Status.Notification.ID)
		try.To1(s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
			ID:       question.Status.Notification.ID,
			ClientID: &agency.ClientID{ID: question.Status.ClientID.ID},
			Ack:      true,
		}))
	}

	return nil
}

func (s *ProofStore) SendProofPresentation(id string) (threadID string, err error) {
	defer err2.Handle(&err)

	try.To2(s.GetProof(id))

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_PRESENT_PROOF,
			Role:   agency.Protocol_RESUMER,
			ID:     id,
		},
		State: agency.ProtocolState_ACK,
	}

	try.To1(s.agent.ProtocolClient.Resume(
		context.TODO(),
		state,
	))

	try.To(s.addProofData(id, &proofData{
		id:          id,
		verifier:    false,
		actualState: StateProofPresentation,
	}))

	return threadID, nil
}

func (s *ProofStore) ProposeProof(connectionID string, attributes []*ProofAttribute) (threadID string, err error) {
	defer err2.Handle(&err)

	log.Printf("Propose proof, conn id: %s, attrs: %v", connectionID, attributes)

	protocol := &agency.Protocol{
		ConnectionID: connectionID,
		TypeID:       agency.Protocol_PRESENT_PROOF,
		Role:         agency.Protocol_ADDRESSEE,
		StartMsg: &agency.Protocol_PresentProof{
			PresentProof: &agency.Protocol_PresentProofMsg{
				AttrFmt: &agency.Protocol_PresentProofMsg_Attributes{
					Attributes: &agency.Protocol_Proof{
						Attributes: attributes,
					},
				},
			},
		},
	}
	res := try.To1(s.agent.Conn.DoStart(context.TODO(), protocol))

	return res.ID, nil
}

func (s *ProofStore) RequestProof(
	connectionID string,
	attributes []*ProofAttribute,
	predicates []*ProofPredicate,
) (threadID string, err error) {
	defer err2.Handle(&err)

	log.Printf("Request proof, conn id: %s, attrs: %v", connectionID, attributes)

	protocol := &agency.Protocol{
		ConnectionID: connectionID,
		TypeID:       agency.Protocol_PRESENT_PROOF,
		Role:         agency.Protocol_INITIATOR,
		StartMsg: &agency.Protocol_PresentProof{
			PresentProof: &agency.Protocol_PresentProofMsg{
				AttrFmt: &agency.Protocol_PresentProofMsg_Attributes{
					Attributes: &agency.Protocol_Proof{
						Attributes: attributes,
					},
				},
				PredFmt: &agency.Protocol_PresentProofMsg_Predicates{
					Predicates: &agency.Protocol_Predicates{
						Predicates: predicates,
					},
				},
			},
		},
	}
	res := try.To1(s.agent.Conn.DoStart(context.TODO(), protocol))

	return res.ID, nil
}

func (s *ProofStore) VerifyPresentation(id string) (err error) {
	defer err2.Handle(&err)

	var question *QuestionHeader
	question, err = s.getProofQuestion(id)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for (err != nil || question.questionID == "") && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Proof not found, waiting for to receive the proof", id)
		time.Sleep(WaitTime)
		question, err = s.getProofQuestion(id)
	}
	try.To(err)

	log.Printf("Accept proof values with the thread id %s, question id %s", question.clientID, question.questionID)
	try.To1(s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
		ID:       question.questionID,
		ClientID: &agency.ClientID{ID: question.clientID},
		Ack:      true,
	}))

	try.To(s.addProofData(id, &proofData{
		id:          id,
		actualState: StateProofDone,
		verifier:    true,
	}))

	return err
}

func (s *ProofStore) addProofData(id string, c *proofData) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if c != nil {
		if data, ok := s.store[id]; ok {
			c.verifier = data.verifier
			c.id = data.id
			if data.questionID != "" {
				c.questionID = data.questionID
			}
			if data.clientID != "" {
				c.clientID = data.clientID
			}
		}

		s.store[id] = *c
		log.Println("Store proof data id", c.id, "state", c.actualState)
		return nil
	}
	return fmt.Errorf("cannot add non-existent proof with id %s", id)
}

func (s *ProofStore) GetProof(id string) (bool, PresentProofState, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if proof, ok := s.store[id]; ok {
		state := proof.actualState
		verifier := proof.verifier
		log.Println("Proof state", id, state)
		return verifier, state, nil
	}
	return false, 0, fmt.Errorf("proof by the id %s not found", id)
}

func (s *ProofStore) getProofQuestion(id string) (*QuestionHeader, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if proof, ok := s.store[id]; ok {
		q := &QuestionHeader{
			questionID: proof.questionID,
			clientID:   proof.clientID,
		}
		return q, nil
	}
	return nil, fmt.Errorf("proof by the id %s not found", id)
}
