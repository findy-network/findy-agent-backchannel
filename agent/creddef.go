package agent

import (
	"context"
	"log"
	"time"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
)

func (a *Agent) CreateCredDef(schemaID, tag string) (id string, err error) {
	defer err2.Handle(&err)

	res := try.To1(a.Client.AgentClient.CreateCredDef(
		context.TODO(),
		&agency.CredDefCreate{
			SchemaID: schemaID,
			Tag:      tag,
		},
	))

	_, err = a.GetCredDef(res.ID)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for err != nil && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Cred def not found, waiting for cred def to be found in ledger", res.ID)
		time.Sleep(WaitTime)
		_, err = a.GetCredDef(res.ID)
	}
	try.To(err)

	id = res.ID
	log.Printf("CreateCredDef: %s", id)

	return
}

func (a *Agent) GetCredDef(credDefID string) (credDefJSON string, err error) {
	defer err2.Handle(&err)

	res := try.To1(a.Client.AgentClient.GetCredDef(
		context.TODO(), &agency.CredDef{
			ID: credDefID,
		},
	))

	credDefJSON = res.Data
	log.Printf("GetCredDef: %v", credDefJSON)

	return
}
