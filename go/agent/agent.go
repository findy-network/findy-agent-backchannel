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

type Agent struct {
	User string
	JWT  string
	Conn client.Conn
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

	authnCmd.UserName = fmt.Sprintf("findy-test-harness-%d", time.Now().Unix())

	myCmd := authnCmd
	myCmd.SubCmd = "register"

	err2.Check(myCmd.Validate())
	_, err := myCmd.Exec(os.Stdout)
	err2.Check(err)
	return &Agent{User: authnCmd.UserName}
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
		"localhost",
		50051,
		[]grpc.DialOption{},
	)
	conn := client.TryAuthOpen(a.JWT, conf)
	a.Conn = conn

}

func (a *Agent) CreateInvitation() (map[string]interface{}, error) {
	sc := agency.NewAgentServiceClient(a.Conn)
	id := uuid.New().String()

	invitation, err := sc.CreateInvitation(
		context.TODO(),
		&agency.InvitationBase{Label: a.User, ID: id},
	)
	if err == nil {
		fmt.Printf("Created invitation\n %s\n", invitation.JSON)
		return map[string]interface{}{"connection_id": id, "invitation": invitation.JSON}, nil
	}
	return nil, err
}
