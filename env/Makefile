AATH_VERSION:=f17f392748ad932d7f4c5c3c3267f3598ef0104e
VON_NETWORK_VERSION:=b041b9a9af79861db0f11664fa7c386d772b845b
EXPECTED_RESULT_AIP10:="3 features passed"
EXPECTED_RESULT_AIP20:="1 feature passed"
INCLUDE_TAGS?=@AcceptanceTest
INCLUDE_TAGS_AIP20?=@T001-RFC0023
AGENT_DEFAULT?=acapy-main
AGENT_BOB?=findy
DATE:=$(shell date +%s)

aath-up: generate clone rm up build-aath
aath-up-check: generate clone rm up build-aath-check
aath-up-check-aip20: generate clone rm up build-aath-check-aip20

generate:
	$(eval DOCKERHOST=$(shell ./dockerhost.sh))
	@cat ./docker-compose.yml.template > ./docker-compose.yml
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST)")}1' ./docker-compose.yml > ./docker-compose.yml.tmp && \
		mv ./docker-compose.yml.tmp ./docker-compose.yml

clone:
	mkdir -p .docker
# enable to run on arm
# -git clone https://github.com/lauravuo/aries-agent-test-harness .docker/aries-agent-test-harness
# cd .docker/aries-agent-test-harness && git checkout arm-support
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness
	cd .docker/aries-agent-test-harness && git reset --hard $(AATH_VERSION)

aath: clone pull-images aath-test-aip10
aath-check: pull-images aath-test-aip10-check
aath-check-aip20: pull-images aath-test-aip20-check

build-aath: clone image aath-test-aip10
build-aath-check: image aath-test-aip10-check
build-aath-check-aip20: image aath-test-aip20-check

pull-images:
	docker pull ghcr.io/findy-network/findy-agent-backchannel/$(AGENT_DEFAULT)-bc:latest
	docker tag ghcr.io/findy-network/findy-agent-backchannel/$(AGENT_DEFAULT)-bc:latest $(AGENT_DEFAULT)-agent-backchannel:latest
	docker pull ghcr.io/findy-network/findy-agent-backchannel/$(AGENT_BOB)-bc:latest
	docker tag ghcr.io/findy-network/findy-agent-backchannel/$(AGENT_BOB)-bc:latest $(AGENT_BOB)-agent-backchannel:latest
	docker pull ghcr.io/findy-network/findy-agent-backchannel/aath:latest
	docker tag ghcr.io/findy-network/findy-agent-backchannel/aath:latest aries-test-harness:latest

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
aath-test-aip10: gen-findy info
	cd .docker/aries-agent-test-harness && \
		./manage run \
			-d $(AGENT_DEFAULT) \
			-b $(AGENT_BOB) \
			-t $(INCLUDE_TAGS) \
			-t @AIP10 \
			-t ~@revocation \
			-t ~@wip \
			-t ~@RFC0025 \
			2>&1 | tee output_aip10.log

aath-test-aip10-check: record-container-logs aath-test-aip10 store-logs
	cat .docker/aries-agent-test-harness/output_aip10.log | grep $(EXPECTED_RESULT_AIP10)

aath-test-check: aath-test-aip10-check

aath-test-aip20: gen-findy info
	cd .docker/aries-agent-test-harness && \
		./manage run \
			-d $(AGENT_DEFAULT) \
			-b $(AGENT_BOB) \
			-t $(INCLUDE_TAGS_AIP20) \
			2>&1 | tee output_aip20.log

aath-test-aip20-check: record-container-logs aath-test-aip20 store-logs
	cat .docker/aries-agent-test-harness/output_aip20.log | grep $(EXPECTED_RESULT_AIP20)

aath-test-check-aip20: aath-test-aip20-check

store-logs:
	-cp .docker/aries-agent-test-harness/output_aip10.log ../tests_output/output_aip10_$(DATE).log
	-cp .docker/aries-agent-test-harness/output_aip20.log ../tests_output/output_aip20_$(DATE).log

record-container-logs:
	mkdir -p ../tests_output
	nohup ./record_logs.sh acme_agent &
	nohup ./record_logs.sh bob_agent &
	nohup ./record_logs.sh faber_agent &
	nohup ./record_logs.sh mallory_agent &

image: pull-images
	cd .. && \
		docker build --build-arg FAB_LOG_INCOMING_REQUESTS=true -t findy-agent-backchannel .

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
# remove build step if on ARM branch
	cd .docker/aries-agent-test-harness/services/von-network/.build/von-network && \
		git fetch origin --depth=1 $(VON_NETWORK_VERSION):version && \
		git checkout version && \
		./manage build
	cd .docker/aries-agent-test-harness && ./manage service start von-network

tails-up:
	-cd .docker/aries-agent-test-harness && ./manage service start indy-tails

resolver-up:
	-cd .docker/aries-agent-test-harness && ./manage service start uniresolver

services-up: von-up tails-up resolver-up

up: services-up gen-cert
	./wait-for-ledger.sh
	mkdir -p conf
	curl http://localhost:9000/genesis > conf/genesis.txt
	docker compose up -d

down:
	docker compose down
	cd .docker/aries-agent-test-harness && ./manage stop && ./manage rm

rm: down
	docker compose rm
	cd .docker/aries-agent-test-harness && ./manage service stop von-network && ./manage service clean von-network
	cd .docker/aries-agent-test-harness && ./manage service stop uniresolver && ./manage service clean uniresolver
	cd .docker/aries-agent-test-harness && ./manage service stop indy-tails && ./manage service clean indy-tails
	rm -rf .data

clean: rm
	rm -rf .docker

