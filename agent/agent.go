package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/findy-network/findy-agent-auth/acator/authn"
	"github.com/findy-network/findy-common-go/agency/client"
	"github.com/findy-network/findy-common-go/agency/client/async"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"google.golang.org/grpc"
)

const AgencyPort = 50051

type Agent struct {
	User           string
	JWT            string
	Conn           client.Conn
	AgentClient    agency.AgentServiceClient
	ProtocolClient agency.ProtocolServiceClient
	AgencyHost     string
	// TODO: create dedicated structure instead of sync.Map
	Invitations      sync.Map
	Connections      sync.Map
	CredentialOffers sync.Map
	Credentials      sync.Map
}

type Credential struct {
	ID        string `json:"referent"`
	CredDefID string `json:"cred_def_id"`
	SchemaID  string `json:"schema_id"`
}

var authnCmd = authn.Cmd{
	SubCmd:   "",
	UserName: "",
	Url:      "http://localhost:8888",
	AAGUID:   "12c85a48-4baf-47bd-b51f-f192871a1511",
	Key:      "15308490f1e4026284594dd08d31291bc8ef2aeac730d0daf6ff87bb92d4336c",
	Counter:  0,
	Token:    "",
	Origin:   "localhost:8888",
}

func Init() *Agent {
	url := os.Getenv("DOCKERHOST")
	if url == "" {
		url = "localhost"
	}
	authnCmd.Url = fmt.Sprintf("http://%s:8888", url)
	authnCmd.UserName = fmt.Sprintf("findy-test-harness-%d", time.Now().UnixNano())

	myCmd := authnCmd
	myCmd.SubCmd = "register"

	err2.Check(myCmd.Validate())
	_, err := myCmd.Exec(os.Stdout)
	err2.Check(err)
	return &Agent{
		User:             authnCmd.UserName,
		AgencyHost:       url,
		Invitations:      sync.Map{},
		Connections:      sync.Map{},
		CredentialOffers: sync.Map{},
		Credentials:      sync.Map{},
	}
}

func (a *Agent) saveConnection(notification *agency.Notification) {
	protocolID := &agency.ProtocolID{
		ID:     notification.ProtocolID,
		TypeID: notification.ProtocolType,
	}
	status, err := a.ProtocolClient.Status(context.TODO(), protocolID)
	err2.Check(err)
	if status.State.State == agency.ProtocolState_OK {
		fmt.Printf("New connection %v\n", status.GetDIDExchange())
		a.Connections.Store(status.GetDIDExchange().ID, status.GetDIDExchange())
	}
}

func (a *Agent) saveCredential(notification *agency.Notification) {
	protocolID := &agency.ProtocolID{
		ID:     notification.ProtocolID,
		TypeID: notification.ProtocolType,
	}
	status, err := a.ProtocolClient.Status(context.TODO(), protocolID)
	err2.Check(err)
	if status.State.State == agency.ProtocolState_OK {
		fmt.Printf("New connection %v\n", status.GetDIDExchange())
		a.Credentials.Store(notification.GetProtocolID(), status.GetIssueCredential())
	}
}

func (a *Agent) handleCredOffer(notification *agency.Notification) {
	a.CredentialOffers.Store(notification.GetProtocolID(), notification)
}

func (a *Agent) Login() {
	myCmd := authnCmd
	myCmd.SubCmd = "login"

	err2.Check(myCmd.Validate())
	r, err := myCmd.Exec(os.Stdout)
	err2.Check(err)

	a.JWT = r.Token

	conf := client.BuildClientConnBase(
		"./env/cert",
		a.AgencyHost,
		AgencyPort,
		[]grpc.DialOption{},
	)

	a.Conn = client.TryAuthOpen(a.JWT, conf)
	a.AgentClient = agency.NewAgentServiceClient(a.Conn)
	a.ProtocolClient = agency.NewProtocolServiceClient(a.Conn)

	ch, err := a.Conn.ListenStatus(context.TODO(), &agency.ClientID{ID: uuid.New().String()})
	err2.Check(err)

	go func() {
		for {
			chRes, ok := <-ch
			if !ok {
				panic("Listening failed")
			}
			notification := chRes.GetNotification()
			fmt.Printf("Received agent notification %v\n", notification)

			if notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
				if notification.GetProtocolType() == agency.Protocol_DIDEXCHANGE {
					a.saveConnection(notification)
				} else if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL {
					a.saveCredential(notification)
				}
			} else if notification.GetProtocolType() == agency.Protocol_ISSUE_CREDENTIAL &&
				notification.GetTypeID() == agency.Notification_PROTOCOL_PAUSED {
				a.handleCredOffer(notification)
			}
		}
	}()
}

func (a *Agent) CreateInvitation() (map[string]interface{}, error) {
	id := uuid.New().String()

	invitation, err := a.AgentClient.CreateInvitation(
		context.TODO(),
		&agency.InvitationBase{Label: a.User, ID: id},
	)
	err2.Check(err)

	fmt.Printf("Created invitation\n %s\n", invitation.JSON)
	var invitationMap map[string]interface{}
	err = json.Unmarshal([]byte(invitation.JSON), &invitationMap)
	err2.Check(err)

	return map[string]interface{}{"connection_id": id, "invitation": invitationMap}, nil
}

func (a *Agent) ReceiveInvitation(invitation map[string]interface{}) (string, error) {
	id := invitation["@id"].(string)
	a.Invitations.Store(id, invitation)
	return id, nil
}

func (a *Agent) GetInvitation(id string) (map[string]interface{}, error) {
	res, ok := a.Invitations.Load(id)
	if !ok {
		panic("Invitation not found")
	}
	return res.(map[string]interface{}), nil
}

func (a *Agent) Connect(invitationID string) (string, error) {
	invitation, _ := a.GetInvitation(invitationID)
	invitationBytes, _ := json.Marshal(invitation)

	pw := async.NewPairwise(a.Conn, "")
	pw.Label = authnCmd.UserName
	_, err := pw.Connection(context.TODO(), string(invitationBytes))
	err2.Check(err)

	return invitationID, nil
}

func (a *Agent) GetConnection(id string) (string, error) {
	_, ok := a.Connections.Load(id)
	if !ok {
		panic("Connection not found")
	}
	return id, nil
}

func (a *Agent) QueryConnection(id string) (string, bool) {
	_, ok := a.Connections.Load(id)
	if ok {
		return id, ok
	}
	return "", ok
}

func (a *Agent) TrustPing(connectionID string) (string, error) {
	_, _ = a.GetConnection(connectionID)

	pw := async.NewPairwise(a.Conn, connectionID)
	_, err := pw.Ping(context.TODO())
	err2.Check(err)

	return connectionID, nil
}

func (a *Agent) GetCredentialOffer(id string) (string, error) {
	res, ok := a.CredentialOffers.Load(id)
	if !ok {
		panic("Offer not found")
	}
	return res.(*agency.Notification).ProtocolID, nil
}

func (a *Agent) RequestCredential(id string) (string, error) {
	res, ok := a.CredentialOffers.Load(id)
	if !ok {
		panic("Offer not found")
	}
	threadID := res.(*agency.Notification).ProtocolID

	state := &agency.ProtocolState{
		ProtocolID: &agency.ProtocolID{
			TypeID: agency.Protocol_ISSUE_CREDENTIAL,
			Role:   agency.Protocol_RESUMER,
			ID:     id,
		},
		State: agency.ProtocolState_ACK,
	}

	_, err := a.ProtocolClient.Resume(
		context.TODO(),
		state,
	)
	err2.Check(err)

	return threadID, nil
}

func (a *Agent) QueryCredential(id string) (string, bool) {
	_, ok := a.Credentials.Load(id)
	if ok {
		return id, ok
	}
	return "", ok
}

func (a *Agent) GetCredential(id string) (*Credential, error) {
	credObj, ok := a.Credentials.Load(id)
	if !ok {
		panic("Credential not found")
	}
	cred := credObj.(*agency.ProtocolStatus_IssueCredentialStatus)
	return &Credential{
		ID:        id,
		CredDefID: cred.CredDefID,
		SchemaID:  cred.SchemaID,
	}, nil
}
