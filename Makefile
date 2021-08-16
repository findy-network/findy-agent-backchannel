
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
		-t findy-test-harness \
		.

drun:
	docker run -it --rm findy-test-harness

scan:
	@./scripts/scan.sh $(ARGS)
