package agent

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/findy-network/findy-agent-auth/acator/authn"
	"github.com/findy-network/findy-common-go/agency/client"
	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/google/uuid"
	"github.com/lainio/err2"
	"google.golang.org/grpc"
)

const AgencyPort = 50051

type AgencyClient struct {
	Conn           client.Conn
	AgentClient    agency.AgentServiceClient
	ProtocolClient agency.ProtocolServiceClient
}

type Agent struct {
	User       string
	JWT        string
	Client     *AgencyClient
	AgencyHost string
	*ConnectionStore
	*CredentialStore
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
		User:       authnCmd.UserName,
		AgencyHost: url,
	}
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

	conn := client.TryAuthOpen(a.JWT, conf)
	a.Client = &AgencyClient{
		Conn:           conn,
		AgentClient:    agency.NewAgentServiceClient(conn),
		ProtocolClient: agency.NewProtocolServiceClient(conn),
	}
	a.ConnectionStore = InitConnections(a.Client, a.User)
	a.CredentialStore = InitCredentials(a.Client)

	ch, err := a.Client.Conn.ListenStatus(context.TODO(), &agency.ClientID{ID: uuid.New().String()})
	err2.Check(err)

	go func() {
		for {
			chRes, ok := <-ch
			if !ok {
				panic("Listening failed")
			}
			notification := chRes.GetNotification()
			fmt.Printf("Received agent notification %v\n", notification)

			a.HandleConnectionNotification(notification)
			a.HandleCredentialNotification(notification)
		}
	}()
}
