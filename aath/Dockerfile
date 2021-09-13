FROM ghcr.io/findy-network/findy-agent:latest AS findy-agent

FROM golang:1.16-buster AS builder

RUN git clone https://github.com/findy-network/findy-agent-auth.git && \
    cd findy-agent-auth && \
    go build -o /go/bin/findy-agent-auth

ADD . findy-agent-backchannel

RUN cd findy-agent-backchannel && \
    go build -o /go/bin/findy-agent-backchannel

FROM ghcr.io/findy-network/findy-base:indy-1.16.ubuntu-18.04

RUN apt-get update && apt-get install curl -y

COPY --from=findy-agent /findy-agent /findy-agent
COPY --from=builder /go/bin/findy-agent-auth /findy-agent-auth
COPY --from=builder /go/bin/findy-agent-backchannel /findy-agent-backchannel

ADD aath/start.sh /
ADD aath/cert /grpc-cert

RUN mkdir -p /data && \
    echo "{}" > /root/findy.json

# auth
ENV FAA_PORT "8888"
ENV FAA_AGENCY_ADDR "localhost"
ENV FAA_AGENCY_PORT "50051"
ENV FAA_AGENCY_ADMIN_ID "findy-root"
ENV FAA_DOMAIN "localhost"
ENV FAA_ORIGIN "localhost:8888"
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
ENV FCLI_AGENCY_POOL_NAME "${FCLI_POOL_NAME}"
ENV FCLI_AGENCY_PSM_DATABASE_FILE "/root/findy.bolt"
ENV FCLI_AGENCY_REGISTER_FILE "/root/findy.json"
ENV FCLI_AGENCY_STEWARD_WALLET_NAME "steward"
ENV FCLI_AGENCY_STEWARD_WALLET_KEY "9C5qFG3grXfU9LodHdMop7CNVb3HtKddjgRc7oK5KhWY"
ENV FCLI_AGENCY_STEWARD_DID "Th7MpTaRZVRYnPiabds81Y"
ENV FCLI_AGENCY_SALT "this is only example"
ENV FCLI_AGENCY_APNS_P12_FILE ""

# backchannel
ENV AGENCY_HOST "localhost"
ENV CERT_PATH "/grpc-cert"

ENTRYPOINT [ "/start.sh" ]