#!/bin/bash

BC_PORT=$2
AGENT_PORT=$(($BC_PORT+1))
AGENT_HOST=$(./ngrok-wait.sh)
SERVER_PORT=$AGENT_PORT

# trim ngrok endpoint, and let agency generate full url
if [[ $AGENT_HOST == https* ]]; then
    export FCLI_AGENCY_HOST_SCHEME="https"
    AGENT_PORT="80"
    prefix="https:\/\/"
    AGENT_HOST=$(echo $AGENT_HOST | sed -e s/^$prefix//)
fi

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

nohup /findy-agent agency start \
    --server-port=$AGENT_PORT \
    --host-scheme=$FCLI_AGENCY_HOST_SCHEME \
    --host-port=$AGENT_PORT \
    --server-port=$SERVER_PORT \
    --host-address=$AGENT_HOST > /logs/core.log &

sleep 1

/findy-agent-backchannel -p $BC_PORT