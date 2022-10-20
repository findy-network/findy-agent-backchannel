build:
	go build ./...

vet:
	go vet ./...

shadow:
	@echo Running govet
	go vet -vettool=$(GOPATH)/bin/shadow ./...
	@echo Govet success

check_fmt:
	$(eval GOFILES = $(shell find . -name '*.go'))
	@gofmt -l $(GOFILES)

lint:
	@golangci-lint run

lint_e:
	@$(GOPATH)/bin/golint ./... | grep -v export | cat

test:
	go test -v -p 1 -failfast ./...

logged_test:
	go test -v -p 1 -failfast ./... -args -logtostderr=true -v=10

test_cov:
	go test -v -p 1 -failfast -coverprofile=c.out ./... && go tool cover -html=c.out

check: check_fmt vet shadow

dbuild:
	docker build \
		-t findy-agent-backchannel \
		.

drun:
	docker run -it --rm findy-agent-backchannel

scan:
	@./scripts/scan.sh $(ARGS)

bundle_test:
	docker build -t findy-aath-bundle -f ./aath/Dockerfile .
	docker run -it --rm \
		-e DOCKERHOST="192.168.65.3" \
		-p 9020-9021:9020-9021 \
		findy-aath-bundle -p 9020

# brew install openapi-generator
# TODO:
# following will do hard reset -> add modified files to ignore list
generate:
	rm -rf ./openapi
	openapi-generator generate -i ./api/openapi.yaml -g go-server -o ./openapi
	-rm ./openapi/go.mod
	-rm ./openapi/Dockerfile
	mv ./openapi/main.go ./
	mv ./openapi/go/* ./openapi/
	rmdir ./openapi/go
	rm -rf ./openapi/api

