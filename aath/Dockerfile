FROM ghcr.io/findy-network/findy-agent:latest AS findy-agent

FROM golang:1.21-bookworm AS builder

RUN git clone https://github.com/findy-network/findy-agent-auth.git && \
  cd findy-agent-auth && \
  go build -o /go/bin/findy-agent-auth

ADD . findy-agent-backchannel

RUN cd findy-agent-backchannel && \
  go install -ldflags "-X 'github.com/findy-network/findy-agent-backchannel/openapi.Version=$(cat ./AGENT_VERSION)'" \
  ./...

FROM ghcr.io/findy-network/findy-base:indy-1.16.ubuntu-22.04

RUN apt-get update && apt-get install curl jq -y

COPY --from=findy-agent /findy-agent /findy-agent
COPY --from=builder /go/bin/findy-agent-auth /findy-agent-auth
COPY --from=builder /go/bin/findy-agent-backchannel /findy-agent-backchannel

ADD aath/start.sh /
ADD aath/ngrok-wait.sh /
ADD aath/cert /grpc-cert

RUN mkdir -p /data && \
  echo "{}" > /root/findy.json

# auth
ENV FAA_PORT "8888"
ENV FAA_AGENCY_ADDR "localhost"
ENV FAA_AGENCY_PORT "50051"
ENV FAA_AGENCY_ADMIN_ID "findy-root"
ENV FAA_DOMAIN "localhost"
ENV FAA_ORIGIN "http://localhost:8888"
ENV FAA_JWT_VERIFICATION_KEY "mySuperSecretKeyLol"
ENV FAA_SEC_KEY "15308490f1e4026284594dd08d31291bc8ef2aeac730d0daf6ff87bb92d4336c"
ENV FAA_LOG_LEVEL "10"
ENV FAA_ENABLE_CORS "false"
ENV FAA_LOCAL_TLS "false"
ENV FAA_TIMEOUT_SECS "60"

# agent
ENV FCLI_LOGGING "-logtostderr=true -v=3"
ENV FCLI_POOL_GENESIS_TXN_FILE "/genesis.txt"
ENV FCLI_POOL_NAME "interop"
ENV FCLI_AGENCY_GRPC_TLS "true"
ENV FCLI_AGENCY_GRPC_CERT_PATH "/grpc-cert"
ENV FCLI_AGENCY_POOL_NAME "FINDY_LEDGER,${FCLI_POOL_NAME},FINDY_MEM_LEDGER,cache"
ENV FCLI_AGENCY_PSM_DATABASE_FILE "/root/findy.bolt"
ENV FCLI_AGENCY_REGISTER_FILE "/root/findy.json"
ENV FCLI_AGENCY_HOST_SCHEME="http"

# backchannel
ENV AGENCY_HOST "localhost"
ENV CERT_PATH "/grpc-cert"

ENTRYPOINT [ "/start.sh" ]
