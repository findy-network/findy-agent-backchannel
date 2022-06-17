package agent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	agency "github.com/findy-network/findy-common-go/grpc/agency/v1"
	"github.com/lainio/err2"
	"github.com/lainio/err2/try"
)

type SchemaStore struct {
	agent   *AgencyClient
	schemas map[string]string
	mtx     sync.RWMutex
}

func InitSchemas(a *AgencyClient) *SchemaStore {
	return &SchemaStore{
		agent:   a,
		schemas: make(map[string]string),
	}
}

func (s *SchemaStore) CreateSchema(name, version string, attributes []string) (id string, err error) {
	defer err2.Return(&err)

	storeID := name + version

	schemaID := ""
	if schemaID, err = s.GetStoredSchema(storeID); err == nil {
		return schemaID, nil
	}

	res := try.To1(s.agent.AgentClient.CreateSchema(
		context.TODO(),
		&agency.SchemaCreate{
			Name:       name,
			Version:    version,
			Attributes: attributes,
		},
	))

	_, err = s.GetSchema(res.ID)
	var totalWaitTime time.Duration
	// TODO: use waitgroups or such
	for err != nil && totalWaitTime < MaxWaitTime {
		totalWaitTime += WaitTime
		log.Println("Schema not found, waiting for schema to be found in ledger", res.ID)
		time.Sleep(WaitTime)
		_, err = s.GetSchema(res.ID)
	}
	try.To(err)

	id = res.ID
	log.Printf("CreateSchema: %s", id)

	try.To1(s.AddStoredSchema(storeID, id))

	return id, nil
}

func (s *SchemaStore) GetSchema(schemaID string) (schemaJSON string, err error) {
	defer err2.Return(&err)

	res := try.To1(s.agent.AgentClient.GetSchema(
		context.TODO(), &agency.Schema{
			ID: schemaID,
		},
	))

	schemaJSON = res.Data
	log.Printf("GetSchema: %v", schemaJSON)

	return
}

func (s *SchemaStore) GetStoredSchema(id string) (string, error) {
	s.mtx.RLock()
	defer s.mtx.RUnlock()
	if schemaID, ok := s.schemas[id]; ok {
		return schemaID, nil
	}
	return "", fmt.Errorf("schema by the store id %s not found", id)
}

func (s *SchemaStore) AddStoredSchema(storeID, id string) (string, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	if id != "" {
		s.schemas[storeID] = id
		return id, nil
	}
	return "", errors.New("cannot add non-existent schema")
}
