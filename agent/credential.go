package agent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
)

type CredentialStatus = agency.ProtocolStatus_IssueCredentialStatus
type CredentialAttribute = agency.Protocol_IssuingAttributes_Attribute
type CredentialProposal = agency.Question_IssueProposeMsg

type CredentialQuestion struct {
	header   QuestionHeader
	proposal *CredentialProposal
}

type Credential struct {
	ID        string `json:"referent"`
	CredDefID string `json:"cred_def_id"`
	SchemaID  string `json:"schema_id"`
}

type CredentialStore struct {
	agent             *AgencyClient
	creds             map[string]*CredentialStatus
	offers            map[string]string
	acceptedProposals map[string]string
	proposals         map[string]*CredentialQuestion
	issuedCreds       map[string]string
	sync.RWMutex
}

func InitCredentials(a *AgencyClient) *CredentialStore {
	return &CredentialStore{
		agent:             a,
		creds:             make(map[string]*CredentialStatus),
		offers:            make(map[string]string),
		acceptedProposals: make(map[string]string),
		proposals:         make(map[string]*CredentialQuestion),
		issuedCreds:       make(map[string]string),
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
				// save cred only if we are holder
				// TODO: role in notification should indicate this
				if _, err = s.GetCredentialOffer(notification.ProtocolID); err == nil {
					cred := status.GetIssueCredential()
					log.Printf("New credential %v\n", cred)
					_, err = s.AddCredential(protocolID.ID, cred)
					err2.Check(err)
				} else {
					_, err = s.AddIssuedCredential(protocolID.ID)
					err2.Check(err)
				}
			}
		}

		// Cred offer received
	} else if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL &&
		notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED {
		_, err = s.AddCredentialOffer(notification.ProtocolID)
		err2.Check(err)
	}
	return nil
}

func (s *CredentialStore) HandleCredentialQuestion(question *agency.Question) (err error) {
	defer err2.Return(&err)
	if question.TypeID == agency.Question_ISSUE_PROPOSE_WAITS {
		_, err := s.AddCredentialProposal(question.Status.Notification.ProtocolID, &CredentialQuestion{
			header: QuestionHeader{
				questionID: question.Status.Notification.ID,
				clientID:   question.Status.ClientID.ID,
			},
			proposal: question.GetIssuePropose(),
		})
		err2.Check(err)
		_, err = s.GetPendingCredentialProposal(question.Status.Notification.ProtocolID)
		if err == nil {
			// proposal is pending, proceed directly
			_, err = s.AcceptCredentialProposal(question.Status.Notification.ProtocolID)
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

	return res.ID, nil
}

func (s *CredentialStore) RequestCredential(id string) (threadID string, err error) {
	defer err2.Return(&err)

	threadID, err = s.GetCredentialOffer(id)
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
	header, err = s.GetCredentialProposal(id)
	err2.Check(err)

	log.Printf("Accept credential proposal with the thread id %s, question id %s", id, header.questionID)
	_, err = s.agent.AgentClient.Give(context.TODO(), &agency.Answer{
		ID:       header.questionID,
		ClientID: &agency.ClientID{ID: header.clientID},
		Ack:      true,
	})
	err2.Check(err)

	return id, nil
}

func (s *CredentialStore) AddCredential(id string, c *CredentialStatus) (*Credential, error) {
	s.Lock()
	defer s.Unlock()
	if c != nil {
		s.creds[id] = c
		res := &Credential{
			ID:        id,
			CredDefID: c.CredDefID,
			SchemaID:  c.SchemaID,
		}
		return res, nil
	}
	return nil, fmt.Errorf("cannot add non-existent credential with id %s", id)
}

func (s *CredentialStore) GetCredential(id string) (*Credential, error) {
	s.RLock()
	defer s.RUnlock()
	if cred, ok := s.creds[id]; ok {
		res := &Credential{
			ID:        id,
			CredDefID: cred.CredDefID,
			SchemaID:  cred.SchemaID,
		}
		return res, nil
	}
	return nil, fmt.Errorf("credential by the id %s not found", id)
}

func (s *CredentialStore) AddCredentialOffer(id string) (string, error) {
	s.Lock()
	defer s.Unlock()
	if id != "" {
		s.offers[id] = id
		return id, nil
	}
	return "", errors.New("cannot add non-existent credential offer")
}

func (s *CredentialStore) GetCredentialOffer(id string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if offer, ok := s.offers[id]; ok {
		return offer, nil
	}
	return "", fmt.Errorf("credential offer by the id %s not found", id)
}

func (s *CredentialStore) AddCredentialProposal(id string, proposal *CredentialQuestion) (string, error) {
	s.Lock()
	defer s.Unlock()
	if proposal != nil {
		s.proposals[id] = proposal
		return id, nil
	}
	return "", errors.New("cannot add non-existent credential proposal")
}

func (s *CredentialStore) GetCredentialProposal(id string) (*QuestionHeader, error) {
	s.RLock()
	defer s.RUnlock()
	if proposal, ok := s.proposals[id]; ok {
		h := &QuestionHeader{
			clientID:   proposal.header.clientID,
			questionID: proposal.header.questionID,
		}
		return h, nil
	}
	return nil, fmt.Errorf("credential proposal by the id %s not found", id)
}

func (s *CredentialStore) AddIssuedCredential(id string) (string, error) {
	s.Lock()
	defer s.Unlock()
	if id != "" {
		s.issuedCreds[id] = id
		return id, nil
	}
	return "", errors.New("cannot add non-existent credential")
}

func (s *CredentialStore) GetIssuedCredential(id string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if cred, ok := s.issuedCreds[id]; ok {
		return cred, nil
	}
	return "", fmt.Errorf("issued credential by the id %s not found", id)
}

func (s *CredentialStore) AddPendingCredentialProposal(id string) (string, error) {
	s.Lock()
	defer s.Unlock()
	if id != "" {
		s.acceptedProposals[id] = id
		return id, nil
	}
	return "", errors.New("cannot add non-existent pending credential proposal")
}

func (s *CredentialStore) GetPendingCredentialProposal(id string) (string, error) {
	s.RLock()
	defer s.RUnlock()
	if cred, ok := s.acceptedProposals[id]; ok {
		return cred, nil
	}
	return "", fmt.Errorf("pending credential proposal by the id %s not found", id)
}
