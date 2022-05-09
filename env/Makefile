AATH_VERSION:=3f08633a3937909a40666e90806dee6bfaea00ab
EXPECTED_RESULT:="3 features passed"
INCLUDE_TAGS?=@AcceptanceTest
AGENT_DEFAULT?=acapy-main
AGENT_BOB?=findy

aath-up: generate clone rm up aath
aath-up-check: generate clone rm up aath-check

generate:
	$(eval DOCKERHOST=$(shell ./dockerhost.sh))
	@cat ./docker-compose.yml.template > ./docker-compose.yml
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST)")}1' ./docker-compose.yml > ./docker-compose.yml.tmp && \
		mv ./docker-compose.yml.tmp ./docker-compose.yml

clone:
	mkdir -p .docker
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness
	cd .docker/aries-agent-test-harness && git reset --hard $(AATH_VERSION)

aath: clone image aath-build aath-test
aath-check: image aath-build aath-test-check

aath-build: 
	cd .docker/aries-agent-test-harness && \
		./manage build -a acapy-main

gen-findy:
	cd .docker/aries-agent-test-harness && \
		mkdir -p aries-backchannels/findy && \
		touch aries-backchannels/findy/Dockerfile.findy && \
		mkdir -p aries-backchannels/findy/.build && \
		rm -rf aries-backchannels/findy/.build/findy.data && \
		cp -R ../../cert aries-backchannels/findy/.build/findy.data

info:
	@printf "\n\n******************************************************\n"
	@printf "\nBob is $(AGENT_BOB), other agents are $(AGENT_DEFAULT)\n"
	@printf "\n******************************************************\n\n"

# exclude tests with '~': WIP and special cases that lack support from acapy/findy
aath-test: gen-findy info
	cd .docker/aries-agent-test-harness && \
		./manage run \
			-d $(AGENT_DEFAULT) \
			-b $(AGENT_BOB) \
			-t $(INCLUDE_TAGS) \
			-t @AIP10 \
			-t ~@revocation \
			-t ~@wip \
			-t ~@RFC0025
	
aath-test-check:  gen-findy info
	cd .docker/aries-agent-test-harness && \
		./manage run \
			-d $(AGENT_DEFAULT) \
			-b $(AGENT_BOB) \
			-t $(INCLUDE_TAGS) \
			-t @AIP10 \
			-t ~@revocation \
			-t ~@wip \
			-t ~@RFC0025 \
			> output.log && \
		cat output.log && \
		cat output.log | grep $(EXPECTED_RESULT)

 # with this image, backchannel will create endorser itself
 # to test the agency capabilities without a steward
image:
	cd .. && docker build --build-arg REGISTER_DID=true -t findy-agent-backchannel .

gen-cert:
	$(eval DOCKERHOST_IP=$(shell docker run --rm --net=host eclipse/che-ip))
	@cat ./cert/client/conf.template > ./cert/client/cert.conf
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST_IP)")}1' ./cert/client/cert.conf > ./cert/client/cert.conf.tmp && \
		mv ./cert/client/cert.conf.tmp ./cert/client/cert.conf
	cd ./cert && ./gen.sh client
	@cat ./cert/server/conf.template > ./cert/server/cert.conf
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST_IP)")}1' ./cert/server/cert.conf > ./cert/server/cert.conf.tmp && \
		mv ./cert/server/cert.conf.tmp ./cert/server/cert.conf
	cd ./cert && ./gen.sh server

von-up:
	-cd .docker/aries-agent-test-harness && ./manage service start von-network
	cd .docker/aries-agent-test-harness/services/von-network/.build/von-network && ./manage build
	cd .docker/aries-agent-test-harness && ./manage service start von-network

tails-up:
	-cd .docker/aries-agent-test-harness && ./manage service start indy-tails
	cd .docker/aries-agent-test-harness/services/indy-tails/.build/indy-tails-server/docker && ./manage build
	cd .docker/aries-agent-test-harness && ./manage service start indy-tails

resolver-up:
	-cd .docker/aries-agent-test-harness && ./manage service start uniresolver

services-up: von-up tails-up resolver-up

up: services-up gen-cert
	./wait-for-ledger.sh
	mkdir -p conf
	curl http://localhost:9000/genesis > conf/genesis.txt
	docker-compose up -d

down:
	docker-compose down
	cd .docker/aries-agent-test-harness && ./manage stop && ./manage rm

rm: down
	docker-compose rm
	cd .docker/aries-agent-test-harness && ./manage service stop von-network && ./manage service clean von-network
	cd .docker/aries-agent-test-harness && ./manage service stop uniresolver && ./manage service clean uniresolver
	cd .docker/aries-agent-test-harness && ./manage service stop indy-tails && ./manage service clean indy-tails
	rm -rf .data

clean: rm
	rm -rf .docker

