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

type Credential struct {
	ID        string `json:"referent"`
	CredDefID string `json:"cred_def_id"`
	SchemaID  string `json:"schema_id"`
}

type CredentialStore struct {
	agent  *AgencyClient
	creds  map[string]CredentialStatus
	offers map[string]string
	sync.RWMutex
}

func InitCredentials(a *AgencyClient) *CredentialStore {
	return &CredentialStore{
		agent:  a,
		creds:  make(map[string]CredentialStatus),
		offers: make(map[string]string),
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
				_, err = s.AddCredential(protocolID.ID, cred)
				err2.Check(err)
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

func (s *CredentialStore) ProposeCredential(connectionID, credDefID string, attributes []*CredentialAttribute) (threadID string, err error) {
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

func (s *CredentialStore) AddCredential(id string, c *CredentialStatus) (*Credential, error) {
	s.Lock()
	defer s.Unlock()
	if c != nil {
		s.creds[id] = *c
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
