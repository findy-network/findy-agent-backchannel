FROM golang:1.16 AS build
WORKDIR /go/src

COPY go.* ./
RUN go mod download

COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o openapi .

FROM scratch AS runtime

COPY --from=build /go/src/openapi ./

EXPOSE "9020-9059"

ADD ./env/cert ./env/cert

ENTRYPOINT ["./openapi"]
