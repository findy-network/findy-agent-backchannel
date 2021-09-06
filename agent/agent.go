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
	authnCmd.UserName = fmt.Sprintf("findy-agent-backchannel-%d", time.Now().UnixNano())

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

func (a *Agent) CreateSchema(name, version string, attributes []string) (id string, err error) {
	defer err2.Return(&err)

	var res *agency.Schema
	res, err = a.Client.AgentClient.CreateSchema(
		context.TODO(),
		&agency.SchemaCreate{
			Name:       name,
			Version:    version,
			Attributes: attributes,
		},
	)
	err2.Check(err)

	id = res.ID
	log.Printf("CreateSchema: %s", id)

	return
}

func (a *Agent) GetSchema(schemaID string) (schemaJSON string, err error) {
	defer err2.Return(&err)

	var res *agency.SchemaData
	res, err = a.Client.AgentClient.GetSchema(
		context.TODO(), &agency.Schema{
			ID: schemaID,
		},
	)
	err2.Check(err)

	schemaJSON = res.Data
	log.Printf("GetSchema: %v", schemaJSON)

	return
}

func (a *Agent) CreateCredDef(schemaID, tag string) (id string, err error) {
	defer err2.Return(&err)

	var res *agency.CredDef
	res, err = a.Client.AgentClient.CreateCredDef(
		context.TODO(),
		&agency.CredDefCreate{
			SchemaID: schemaID,
			Tag:      tag,
		},
	)
	err2.Check(err)

	id = res.ID
	log.Printf("CreateCredDef: %s", id)

	return
}

func (a *Agent) GetCredDef(credDefID string) (credDefJSON string, err error) {
	defer err2.Return(&err)

	var res *agency.CredDefData
	res, err = a.Client.AgentClient.GetCredDef(
		context.TODO(), &agency.CredDef{
			ID: credDefID,
		},
	)
	err2.Check(err)

	credDefJSON = res.Data
	log.Printf("GetCredDef: %v", credDefJSON)

	return
}
