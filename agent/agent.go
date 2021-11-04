package agent

import (
	"context"
	"fmt"
	"log"
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
const WaitTime = time.Millisecond * 100
const MaxWaitTime = time.Minute

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
	*ProofStore
	*SchemaStore
}

type QuestionHeader struct {
	questionID string
	clientID   string
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
	url := os.Getenv("AGENCY_HOST")
	if url == "" {
		url = os.Getenv("DOCKERHOST")
	}
	publicDIDSeed := ""
	if os.Getenv("REGISTER_DID") == "true" {
		publicDIDSeed = registerDID()
	}
	authnCmd.Url = fmt.Sprintf("http://%s:8888", url)
	authnCmd.UserName = fmt.Sprintf("findy-agent-backchannel-%d", time.Now().UnixNano())
	authnCmd.PublicDIDSeed = publicDIDSeed
	log.Printf("Auth url %s, origin %s, user %s", authnCmd.Url, authnCmd.Origin, authnCmd.UserName)

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
		os.Getenv("CERT_PATH"),
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
	a.ProofStore = InitProofs(a.Client)
	a.SchemaStore = InitSchemas(a.Client)

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

			_ = a.HandleConnectionNotification(notification)
			_ = a.HandleCredentialNotification(notification)
			_ = a.HandleProofNotification(notification)
		}
	}()

	questionCh, err := a.Client.Conn.Wait(context.TODO(), &agency.ClientID{ID: uuid.New().String()})
	err2.Check(err)

	go func() {
		for {
			chRes, ok := <-questionCh
			if !ok {
				panic("Waiting failed")
			}
			fmt.Printf("Received question %v\n", chRes)

			_ = a.HandleCredentialQuestion(chRes)
			_ = a.HandleProofQuestion(chRes)
		}
	}()
}
