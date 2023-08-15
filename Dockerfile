FROM golang:1.21 AS build

WORKDIR /go/src

COPY go.* ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o backchannel .

FROM scratch AS runtime

ARG FAB_LOG_INCOMING_REQUESTS=false

COPY --from=build /go/src/backchannel ./

EXPOSE "9020-9059"

ENV CERT_PATH "/data-mount"
ENV FAB_LOG_INCOMING_REQUESTS=${FAB_LOG_INCOMING_REQUESTS}

ENTRYPOINT ["./backchannel"]
