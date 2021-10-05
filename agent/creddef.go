package agent

import (
	"context"
	"log"
	"time"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
)

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

	_, err = a.GetCredDef(res.ID)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for err != nil && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Cred def not found, waiting for cred def to be found in ledger", res.ID)
		time.Sleep(WaitTime)
		_, err = a.GetCredDef(res.ID)
	}
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
