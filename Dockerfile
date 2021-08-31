FROM golang:1.16 AS build
WORKDIR /go/src

COPY go.* ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0

RUN go build -a -installsuffix cgo -o backchannel .

FROM scratch AS runtime

COPY --from=build /go/src/backchannel ./

EXPOSE "9020-9059"

ADD ./env/cert ./env/cert

ENTRYPOINT ["./backchannel"]
