package agent

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
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
	id            string
	questionID    string
	clientID      string
	actualState   PresentProofState
	reportedState PresentProofState
	verifier      bool
}

type Proof struct {
	ID         string `json:"referent"`
	ProofDefID string `json:"proof_def_id"`
	SchemaID   string `json:"schema_id"`
}

type ProofStore struct {
	agent *AgencyClient
	store map[string]proofData
	sync.RWMutex
}

func InitProofs(a *AgencyClient) *ProofStore {
	return &ProofStore{
		agent: a,
		store: make(map[string]proofData),
	}
}

func (s *ProofStore) HandleProofNotification(notification *agency.Notification) (err error) {
	defer err2.Return(&err)

	// Proof success
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}

			var status *agency.ProtocolStatus
			status, err = s.agent.ProtocolClient.Status(context.TODO(), protocolID)
			err2.Check(err)

			if status.State.State == agency.ProtocolState_OK {
				proof := status.GetPresentProof()
				log.Printf("Proof ready %v\n", proof)

				// TODO: role in notification should indicate if we are holder or not
				var verifier bool
				verifier, _, err = s.GetProof(notification.ProtocolID)
				err2.Check(err)

				data := &proofData{
					id:          protocolID.ID,
					verifier:    verifier,
					actualState: StateProofDone,
				}
				err = s.addProofData(protocolID.ID, data)
				err2.Check(err)
			}
		}

		// Proof request received
	} else if notification.GetProtocolType() == agency.Protocol_PRESENT_PROOF &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED {
		data := &proofData{
			id:          notification.ProtocolID,
			verifier:    false,
			actualState: StateProofRequest,
		}
		err = s.addProofData(notification.ProtocolID, data)
		err2.Check(err)
	}
	return nil
}

func (s *ProofStore) HandleProofQuestion(question *agency.Question) (err error) {
	defer err2.Return(&err)

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

		err := s.addProofData(data.id, data)
		err2.Check(err)
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

		err := s.addProofData(data.id, data)
		err2.Check(err)

		// just accept proof propose directly
		log.Printf("Accept proof proposal with the thread id %s, question id %s", question.Status.ClientID.ID, question.Status.Notification.ID)
		_, err = s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
			ID:       question.Status.Notification.ID,
			ClientID: &agency.ClientID{ID: question.Status.ClientID.ID},
			Ack:      true,
		})
		err2.Check(err)
	}

	return nil
}

func (s *ProofStore) SendProofPresentation(id string) (threadID string, err error) {
	defer err2.Return(&err)

	_, _, err = s.GetProof(id)
	err2.Check(err)

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_PRESENT_PROOF,
			Role:   agency.Protocol_RESUMER,
			ID:     id,
		},
		State: agency.ProtocolState_ACK,
	}

	_, err = s.agent.ProtocolClient.Resume(
		context.TODO(),
		state,
	)
	err2.Check(err)

	err = s.addProofData(id, &proofData{
		id:          id,
		verifier:    false,
		actualState: StateProofPresentation,
	})
	err2.Check(err)

	return threadID, nil
}

func (s *ProofStore) ProposeProof(connectionID string, attributes []*ProofAttribute) (threadID string, err error) {
	defer err2.Return(&err)

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
	res, err := s.agent.Conn.DoStart(context.TODO(), protocol)
	err2.Check(err)

	return res.ID, nil
}

func (s *ProofStore) RequestProof(
	connectionID string,
	attributes []*ProofAttribute,
	predicates []*ProofPredicate,
) (threadID string, err error) {
	defer err2.Return(&err)

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
	res, err := s.agent.Conn.DoStart(context.TODO(), protocol)
	err2.Check(err)

	return res.ID, nil
}

func (s *ProofStore) VerifyPresentation(id string) (err error) {
	defer err2.Return(&err)

	var question *QuestionHeader
	question, err = s.getProofQuestion(id)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for err != nil && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Proof not found, waiting for to receive the proof", id)
		time.Sleep(WaitTime)
		question, err = s.getProofQuestion(id)
	}
	err2.Check(err)

	log.Printf("Accept proof values with the thread id %s, question id %s", question.clientID, question.questionID)
	_, err = s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
		ID:       question.questionID,
		ClientID: &agency.ClientID{ID: question.clientID},
		Ack:      true,
	})
	err2.Check(err)

	err = s.doAddProofData(id, &proofData{
		id:          id,
		actualState: StateProofDone,
		verifier:    true,
	}, true)
	err2.Check(err)

	return err
}

func (s *ProofStore) addProofData(id string, c *proofData) error {
	return s.doAddProofData(id, c, false)
}

func (s *ProofStore) doAddProofData(id string, c *proofData, resetReported bool) error {
	s.Lock()
	defer s.Unlock()
	if c != nil {
		reportedState := c.actualState
		if !resetReported {
			if data, ok := s.store[id]; ok {
				reportedState = data.reportedState
			}
		}
		c.reportedState = reportedState
		s.store[id] = *c
		log.Println("Store proof data id", c.id, "state", c.actualState, "reported", c.reportedState)
		return nil
	}
	return fmt.Errorf("cannot add non-existent proof with id %s", id)
}

func (s *ProofStore) GetProof(id string) (bool, PresentProofState, error) {
	s.Lock()
	defer s.Unlock()
	if proof, ok := s.store[id]; ok {
		state := proof.actualState
		verifier := proof.verifier
		/*// we do not get all protocol notifications from agency so simulate here
		// "step-by-step"-functionality
		if proof.actualState > state || state == StateProofDone-1 {
			proof.reportedState++
		}
		s.store[id] = proof
		log.Println("Update reported proof data state", proof.id, "state", proof.actualState, "reported", proof.reportedState)*/
		log.Println("Proof state", id, state)
		return verifier, state, nil
	}
	return false, 0, fmt.Errorf("proof by the id %s not found", id)
}

func (s *ProofStore) getProofQuestion(id string) (*QuestionHeader, error) {
	s.Lock()
	defer s.Unlock()
	if proof, ok := s.store[id]; ok {
		q := &QuestionHeader{
			questionID: proof.questionID,
			clientID:   proof.clientID,
		}
		return q, nil
	}
	return nil, fmt.Errorf("proof by the id %s not found", id)
}
