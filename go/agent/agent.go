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

type Agent struct {
	User           string
	JWT            string
	Conn           client.Conn
	AgentClient    agency.AgentServiceClient
	ProtocolClient agency.ProtocolServiceClient
	AgencyHost     string
	Invitations    sync.Map
	Connections    sync.Map
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

	//authnCmd.Origin = fmt.Sprintf("//%s:8888", url)

	authnCmd.UserName = fmt.Sprintf("findy-test-harness-%d", time.Now().UnixNano())

	myCmd := authnCmd
	myCmd.SubCmd = "register"

	err2.Check(myCmd.Validate())
	_, err := myCmd.Exec(os.Stdout)
	err2.Check(err)
	return &Agent{User: authnCmd.UserName, AgencyHost: url, Invitations: sync.Map{}}
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
		50052,
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
			if notification.GetProtocolType() == agency.Protocol_DIDEXCHANGE &&
				notification.GetTypeID() == agency.Notification_STATUS_UPDATE {
				protocolID := &agency.ProtocolID{
					ID:     notification.ProtocolID,
					TypeID: notification.ProtocolType,
				}
				status, err := a.ProtocolClient.Status(context.TODO(), protocolID)
				err2.Check(err)
				if status.State.State == agency.ProtocolState_OK {
					a.Connections.Store(status.GetDIDExchange().ID, status.GetDIDExchange())
				}
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
	json.Unmarshal([]byte(invitation.JSON), &invitationMap)
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

func (a *Agent) Connect(invitationId string) (string, error) {
	invitation, _ := a.GetInvitation(invitationId)
	invitationBytes, _ := json.Marshal(invitation)

	pw := async.NewPairwise(a.Conn, "")
	_, err := pw.Connection(context.TODO(), string(invitationBytes))
	err2.Check(err)

	return invitationId, nil
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

func (a *Agent) TrustPing(connectionId string) (string, error) {
	_, _ = a.GetConnection(connectionId)

	pw := async.NewPairwise(a.Conn, connectionId)
	_, err := pw.Ping(context.TODO())
	err2.Check(err)

	return connectionId, nil
}
