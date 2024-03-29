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
	FAILURE    IssueCredentialState = 6
)

func (e IssueCredentialState) String() string {
	switch e {
	case PROPOSAL:
		return "PROPOSAL"
	case OFFER:
		return "OFFER"
	case REQUEST:
		return "REQUEST"
	case CREDENTIAL:
		return "CREDENTIAL"
	case DONE:
		return "DONE"
	case FAILURE:
		return "FAILURE"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

type credData struct {
	id          string
	questionID  string
	clientID    string
	actualState IssueCredentialState
	issuer      bool
	credDefID   string
	schemaID    string
}

type CredentialStore struct {
	agent *AgencyClient
	store map[string]credData
	mtx   sync.RWMutex
}

func InitCredentials(a *AgencyClient) *CredentialStore {
	return &CredentialStore{
		agent: a,
		store: make(map[string]credData),
	}
}

func (s *CredentialStore) HandleCredentialNotification(notification *agency.Notification) (err error) {
	defer err2.Handle(&err)

	// Cred issued
	if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
		if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL {
			protocolID := &agency.ProtocolID{
				ID:     notification.ProtocolID,
				TypeID: notification.ProtocolType,
			}

			status := try.To1(s.agent.ProtocolClient.Status(context.TODO(), protocolID))

			if status.State.State == agency.ProtocolState_OK {
				cred := status.GetIssueCredential()
				log.Printf("New credential %v\n", cred)

				// TODO: role in notification should indicate if we are holder or not
				issuer, _ := try.To2(s.GetCredential(notification.ProtocolID))

				state := CREDENTIAL
				if issuer {
					state = REQUEST
				}

				data := &credData{
					id:          protocolID.ID,
					issuer:      issuer,
					actualState: state,
					credDefID:   cred.CredDefID,
					schemaID:    cred.SchemaID,
				}
				try.To(s.addCredData(protocolID.ID, data))
			} else if status.State.State == agency.ProtocolState_ERR {
				log.Printf("Credential status ERR! %v\n", status.State)
				try.To(s.addCredData(protocolID.ID, &credData{
					id:          protocolID.ID,
					actualState: FAILURE,
				}))
			}
		}

		// Cred offer received
	} else if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED &&
		notification.GetRole() == agency.Protocol_ADDRESSEE {
		data := &credData{
			id:          notification.ProtocolID,
			issuer:      false,
			actualState: OFFER,
		}
		try.To(s.addCredData(notification.ProtocolID, data))
	}
	return nil
}

func (s *CredentialStore) HandleCredentialQuestion(question *agency.Question) (err error) {
	defer err2.Handle(&err)
	if question.TypeID == agency.Question_ISSUE_PROPOSE_WAITS {
		data := &credData{
			id:          question.Status.Notification.ProtocolID,
			questionID:  question.Status.Notification.ID,
			clientID:    question.Status.ClientID.ID,
			issuer:      true,
			actualState: REQUEST,
		}
		try.To(s.addCredData(data.id, data))
	}
	return nil
}

func (s *CredentialStore) ProposeCredential(
	connectionID, credDefID string,
	attributes []*CredentialAttribute,
) (threadID string, err error) {
	defer err2.Handle(&err)

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
	res := try.To1(s.agent.Conn.DoStart(context.TODO(), protocol))

	return res.ID, nil
}

func (s *CredentialStore) OfferCredential(
	connectionID, credDefID string,
	attributes []*CredentialAttribute,
) (threadID string, err error) {
	defer err2.Handle(&err)

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
	res := try.To1(s.agent.Conn.DoStart(context.TODO(), protocol))

	try.To(s.addCredData(res.GetID(), &credData{
		id:          res.GetID(),
		actualState: REQUEST,
		issuer:      true,
	}))

	return res.ID, nil
}

func (s *CredentialStore) RequestCredential(id string) (threadID string, err error) {
	defer err2.Handle(&err)

	try.To2(s.GetCredential(id))

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_ISSUE_CREDENTIAL,
			Role:   agency.Protocol_RESUMER,
			ID:     id,
		},
		State: agency.ProtocolState_ACK,
	}

	try.To1(s.agent.ProtocolClient.Resume(
		context.TODO(),
		state,
	))

	return threadID, nil
}

func (s *CredentialStore) AcceptCredentialProposal(id string) (threadID string, err error) {
	defer err2.Handle(&err)

	var header *QuestionHeader
	header, err = s.getCredentialQuestion(id)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for (err != nil || header.questionID == "") && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Credential not found, waiting for to receive the credential", id)
		time.Sleep(WaitTime)
		header, err = s.getCredentialQuestion(id)
	}
	try.To(err)

	log.Printf("Accept credential proposal with the thread id %s, question id %s", id, header.questionID)
	try.To1(s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
		ID:       header.questionID,
		ClientID: &agency.ClientID{ID: header.clientID},
		Ack:      true,
	}))

	try.To(s.addCredData(id, &credData{
		id:          id,
		actualState: REQUEST,
		issuer:      true,
	}))

	return id, nil
}

func (s *CredentialStore) IssueCredential(id string) (err error) {
	defer err2.Handle(&err)

	_, state := try.To2(s.GetCredential(id))

	if state == REQUEST {
		try.To(s.addCredData(id, &credData{
			id:          id,
			actualState: CREDENTIAL,
			issuer:      false,
		}))
	} else {
		try.To(fmt.Errorf("unable to issue credential %s with state %s", id, state))
	}

	return err
}

func (s *CredentialStore) ReceiveCredential(id string) (err error) {
	defer err2.Handle(&err)

	var state IssueCredentialState
	_, state, err = s.GetCredential(id)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for (err != nil || (state != CREDENTIAL && state != FAILURE)) && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Credential not found, waiting for to receive the credential", id)
		time.Sleep(WaitTime)
		_, state, err = s.GetCredential(id)
	}
	try.To(err)
	if state != CREDENTIAL {
		try.To(fmt.Errorf("Credential %s not received", id))
	}

	try.To(s.addCredData(id, &credData{
		id:          id,
		actualState: DONE,
		issuer:      false,
	}))

	return err
}

func (s *CredentialStore) addCredData(id string, c *credData) error {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if c != nil {
		if data, ok := s.store[id]; ok {
			c.issuer = data.issuer
			c.id = data.id
			if data.questionID != "" {
				c.questionID = data.questionID
			}
			if data.clientID != "" {
				c.clientID = data.clientID
			}
			if data.schemaID != "" {
				c.schemaID = data.schemaID
			}
			if data.credDefID != "" {
				c.credDefID = data.credDefID
			}
		}
		s.store[id] = *c
		log.Println("Store cred data id", c.id, "state", c.actualState)
		return nil
	}
	return fmt.Errorf("cannot add non-existent credential with id %s", id)
}

func (s *CredentialStore) GetCredential(id string) (bool, IssueCredentialState, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if cred, ok := s.store[id]; ok {
		state := cred.actualState
		issuer := cred.issuer
		log.Println("Credential state", id, state)
		return issuer, state, nil
	}
	return false, 0, fmt.Errorf("credential by the id %s not found", id)
}

func (s *CredentialStore) getCredentialQuestion(id string) (*QuestionHeader, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
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
	s.mtx.Lock()
	defer s.mtx.Unlock()
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
