FROM golang:1.16-alpine3.13

WORKDIR /work

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o /go/bin/findy-template-go

FROM alpine:3.13

COPY --from=0 /go/bin/findy-template-go /findy-template-go

RUN echo '/findy-template-go' > /start.sh && chmod a+x /start.sh

ENTRYPOINT ["/bin/sh", "-c", "/start.sh"]
