package agent

import (
	"context"
	"fmt"
	"log"
	"sync"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
)

type CredentialStatus = agency.ProtocolStatus_IssueCredentialStatus
type CredentialAttribute = agency.Protocol_IssuingAttributes_Attribute
type CredentialProposal = agency.Question_IssueProposeMsg

type Credential struct {
	ID        string `json:"referent"`
	CredDefID string `json:"cred_def_id"`
	SchemaID  string `json:"schema_id"`
}

type IssueCredentialState int

const (
	PROPOSAL   IssueCredentialState = 1
	OFFER      IssueCredentialState = 2
	REQUEST    IssueCredentialState = 3
	CREDENTIAL IssueCredentialState = 4
	DONE       IssueCredentialState = 5
)

type credData struct {
	id            string
	questionID    string
	clientID      string
	actualState   IssueCredentialState
	reportedState IssueCredentialState
	issuer        bool
	credDefID     string
	schemaID      string
}

type CredentialStore struct {
	agent *AgencyClient
	store map[string]credData
	sync.RWMutex
}

func InitCredentials(a *AgencyClient) *CredentialStore {
	return &CredentialStore{
		agent: a,
		store: make(map[string]credData),
	}
}

func (s *CredentialStore) HandleCredentialNotification(notification *agency.Notification) (err error) {
	defer err2.Return(&err)

	// Cred issued
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}

			var status *agency.ProtocolStatus
			status, err = s.agent.ProtocolClient.Status(context.TODO(), protocolID)
			err2.Check(err)

			if status.State.State == agency.ProtocolState_OK {
				cred := status.GetIssueCredential()
				log.Printf("New credential %v\n", cred)

				// TODO: role in notification should indicate if we are holder or not
				issuer, _, err := s.GetCredential(notification.ProtocolID)
				err2.Check(err)

				data := &credData{
					id:          protocolID.ID,
					issuer:      issuer,
					actualState: CREDENTIAL,
					credDefID:   cred.CredDefID,
					schemaID:    cred.SchemaID,
				}
				err = s.addCredData(protocolID.ID, data)
				err2.Check(err)
			}
		}

		// Cred offer received
	} else if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED {
		data := &credData{
			id:          notification.ProtocolID,
			issuer:      false,
			actualState: OFFER,
		}
		err = s.addCredData(notification.ProtocolID, data)
		err2.Check(err)
	}
	return nil
}

func (s *CredentialStore) HandleCredentialQuestion(question *agency.Question) (err error) {
	defer err2.Return(&err)
	if question.TypeID == agency.Question_ISSUE_PROPOSE_WAITS {

		data := &credData{
			id:          question.Status.Notification.ProtocolID,
			questionID:  question.Status.Notification.ID,
			clientID:    question.Status.ClientID.ID,
			issuer:      false,
			actualState: REQUEST,
		}

		_, state, err := s.GetCredential(question.Status.Notification.ProtocolID)
		if err == nil && state == OFFER {
			err := s.addCredData(data.id, data)
			err2.Check(err)

			_, err = s.AcceptCredentialProposal(question.Status.Notification.ProtocolID)
			err2.Check(err)
		} else {
			data.actualState = PROPOSAL
			err := s.addCredData(data.id, data)
			err2.Check(err)
		}
	}
	return nil
}

func (s *CredentialStore) ProposeCredential(
	connectionID, credDefID string,
	attributes []*CredentialAttribute,
) (threadID string, err error) {
	defer err2.Return(&err)

	log.Printf("Propose credential, conn id: %s, credDefID: %s, attrs: %v", connectionID, credDefID, attributes)

	protocol := &agency.Protocol{
		ConnectionID: connectionID,
		TypeID:       agency.Protocol_ISSUE_CREDENTIAL,
		Role:         agency.Protocol_ADDRESSEE,
		StartMsg: &agency.Protocol_IssueCredential{
			IssueCredential: &agency.Protocol_IssueCredentialMsg{
				CredDefID: credDefID,
				AttrFmt: &agency.Protocol_IssueCredentialMsg_Attributes{
					Attributes: &agency.Protocol_IssuingAttributes{
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

func (s *CredentialStore) OfferCredential(
	connectionID, credDefID string,
	attributes []*CredentialAttribute,
) (threadID string, err error) {
	defer err2.Return(&err)

	log.Printf("Offer credential, conn id: %s, credDefID: %s, attrs: %v", connectionID, credDefID, attributes)

	protocol := &agency.Protocol{
		ConnectionID: connectionID,
		TypeID:       agency.Protocol_ISSUE_CREDENTIAL,
		Role:         agency.Protocol_INITIATOR,
		StartMsg: &agency.Protocol_IssueCredential{
			IssueCredential: &agency.Protocol_IssueCredentialMsg{
				CredDefID: credDefID,
				AttrFmt: &agency.Protocol_IssueCredentialMsg_Attributes{
					Attributes: &agency.Protocol_IssuingAttributes{
						Attributes: attributes,
					},
				},
			},
		},
	}
	res, err := s.agent.Conn.DoStart(context.TODO(), protocol)
	err2.Check(err)

	err = s.addCredData(res.GetID(), &credData{
		id:          res.GetID(),
		actualState: REQUEST,
		issuer:      true,
	})
	err2.Check(err)

	return res.ID, nil
}

func (s *CredentialStore) RequestCredential(id string) (threadID string, err error) {
	defer err2.Return(&err)

	_, _, err = s.GetCredential(id)
	err2.Check(err)

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_ISSUE_CREDENTIAL,
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

	return threadID, nil
}

func (s *CredentialStore) AcceptCredentialProposal(id string) (threadID string, err error) {
	defer err2.Return(&err)

	var header *QuestionHeader
	header, err = s.GetCredentialQuestion(id)
	if err == nil {
		log.Printf("Accept credential proposal with the thread id %s, question id %s", id, header.questionID)
		_, err = s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
			ID:       header.questionID,
			ClientID: &agency.ClientID{ID: header.clientID},
			Ack:      true,
		})
		err2.Check(err)
	}

	err = s.addCredData(id, &credData{
		id:          id,
		actualState: REQUEST,
		issuer:      true,
	})
	err2.Check(err)

	return id, nil
}

func (s *CredentialStore) IssueCredential(id string) (err error) {
	defer err2.Return(&err)
	err = s.addCredData(id, &credData{
		id:          id,
		actualState: CREDENTIAL,
		issuer:      true,
	})
	err2.Check(err)

	return nil
}

func (s *CredentialStore) addCredData(id string, c *credData) error {
	s.Lock()
	defer s.Unlock()
	if c != nil {
		reportedState := c.actualState
		if data, ok := s.store[id]; ok {
			reportedState = data.reportedState
		}
		c.reportedState = reportedState
		s.store[id] = *c
		return nil
	}
	return fmt.Errorf("cannot add non-existent credential with id %s", id)
}

func (s *CredentialStore) GetCredential(id string) (bool, IssueCredentialState, error) {
	s.Lock()
	defer s.Unlock()
	if cred, ok := s.store[id]; ok {
		state := cred.reportedState
		issuer := cred.issuer
		// we do not get all protocol notifications from agency so simulate here
		// "step-by-step"-functionality
		if cred.actualState > state || state == DONE-1 {
			state++
		}
		cred.reportedState = state
		s.store[id] = cred
		return issuer, state, nil
	}
	return false, 0, fmt.Errorf("credential by the id %s not found", id)
}

func (s *CredentialStore) GetCredentialQuestion(id string) (*QuestionHeader, error) {
	s.Lock()
	defer s.Unlock()
	if cred, ok := s.store[id]; ok {
		q := &QuestionHeader{
			questionID: cred.questionID,
			clientID:   cred.clientID,
		}
		return q, nil
	}
	return nil, fmt.Errorf("credential by the id %s not found", id)
}

func (s *CredentialStore) GetCredentialContent(id string) (*Credential, error) {
	s.Lock()
	defer s.Unlock()
	if cred, ok := s.store[id]; ok {
		c := &Credential{
			ID:        id,
			SchemaID:  cred.schemaID,
			CredDefID: cred.credDefID,
		}
		return c, nil
	}
	return nil, fmt.Errorf("credential by the id %s not found", id)
}
