#!/bin/bash

BC_PORT=$2
AGENT_PORT=$(($BC_PORT+1))

echo "Starting backchannel to port $BC_PORT and agent to port $AGENT_PORT"

echo "Fetching genesis txn from $LEDGER_URL/genesis"

curl "${LEDGER_URL}/genesis" > /genesis.txt

mkdir -p logs

nohup /findy-agent-auth \
    --port $FAA_PORT \
    --agency $FAA_AGENCY_ADDR \
    --gport $FAA_AGENCY_PORT \
    --admin $FAA_AGENCY_ADMIN_ID \
    --domain $FAA_DOMAIN \
    --origin $FAA_ORIGIN \
    --sec-file "/data/fido-enclave.bolt" \
    --sec-key $FAA_SEC_KEY \
    --cert-path /grpc-cert \
    --logging "-logtostderr=true -v=$FAA_LOG_LEVEL" \
    --cors=$FAA_ENABLE_CORS \
    --local-tls=$FAA_LOCAL_TLS \
    --jwt-secret $FAA_JWT_VERIFICATION_KEY \
    --timeout $FAA_TIMEOUT_SECS > /logs/auth.log &


/findy-agent ledger pool create \
    --name ${FCLI_POOL_NAME} \
    --genesis-txn-file /genesis.txt
/findy-agent ledger steward create \
    --pool-name ${FCLI_POOL_NAME} \
    --seed 000000000000000000000000Steward1 \
    --wallet-name ${FCLI_AGENCY_STEWARD_WALLET_NAME} \
    --wallet-key ${FCLI_AGENCY_STEWARD_WALLET_KEY}

nohup /findy-agent agency start \
    --server-port=$AGENT_PORT \
    --host-port=$AGENT_PORT \
    --host-address=$DOCKERHOST \
    --grpc=true > /logs/core.log &

sleep 1

/findy-agent-backchannel -p $BC_PORT