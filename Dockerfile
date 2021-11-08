FROM golang:1.16 AS build

WORKDIR /go/src

COPY go.* ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o backchannel .

FROM scratch AS runtime

ARG REGISTER_DID="false"

COPY --from=build /go/src/backchannel ./

EXPOSE "9020-9059"

ENV CERT_PATH "/data-mount"
ENV REGISTER_DID "${REGISTER_DID}"

ENTRYPOINT ["./backchannel"]
