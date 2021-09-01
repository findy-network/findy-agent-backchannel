VON_NETWORK_VERSION:=37fa43673678b64c0f6080c5e2b50c7ea5354cb4
AATH_VERSION:=4aa61a516816cdc52c368c74d5eb3965ec87c18f
EXPECTED_RESULT:="3 features passed"
TEST_TAGS:="@AIP10 -t ~@wip -t ~@T006-RFC0160 -t ~@T003-RFC0160 -t ~@T004-RFC0160"

aath-up: generate clone rm up aath
aath-check: generate clone rm up aath-check

generate:
	$(eval DOCKERHOST=$(shell docker run --rm --net=host eclipse/che-ip))
	@cat ./docker-compose.yml.template > ./docker-compose.yml
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST)")}1' ./docker-compose.yml > ./docker-compose.yml.tmp && \
		mv ./docker-compose.yml.tmp ./docker-compose.yml

clone:
	mkdir -p .docker
	-git clone https://github.com/bcgov/von-network .docker/von-network && \
		cd .docker/von-network && git checkout $(VON_NETWORK_VERSION)
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness && \
		cd .docker/aries-agent-test-harness && git checkout $(AATH_VERSION)

von-up:
	cd .docker/von-network && ./manage build && ./manage start

aath: image aath-build aath-test
aath-check: image aath-build aath-test-check

aath-build: 
	cd .docker/aries-agent-test-harness && \
		./manage build -a acapy

aath-test: 
	cd .docker/aries-agent-test-harness && \
		mkdir -p aries-backchannels/findy && \
		touch aries-backchannels/findy/Dockerfile.findy && \
		./manage run -d acapy -b findy -t $(TEST_TAGS)
	
aath-test-check: 
	cd .docker/aries-agent-test-harness && \
		mkdir -p aries-backchannels/findy && \
		touch aries-backchannels/findy/Dockerfile.findy && \
		./manage run -d acapy -b findy -t $(TEST_TAGS) > output.log && \
		cat output.log && \
		cat output.log | grep $(EXPECTED_RESULT)

image:
	cd .. && docker build -t findy-agent-backchannel .

gen-cert:
	$(eval DOCKERHOST=$(shell docker run --rm --net=host eclipse/che-ip))
	@cat ./cert/client/conf.template > ./cert/client/cert.conf
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST)")}1' ./cert/client/cert.conf > ./cert/client/cert.conf.tmp && \
		mv ./cert/client/cert.conf.tmp ./cert/client/cert.conf
	cd ./cert && ./gen.sh client
	@cat ./cert/server/conf.template > ./cert/server/cert.conf
	awk '{sub("<IP_ADDRESS>","$(DOCKERHOST)")}1' ./cert/server/cert.conf > ./cert/server/cert.conf.tmp && \
		mv ./cert/server/cert.conf.tmp ./cert/server/cert.conf
	cd ./cert && ./gen.sh server

up: von-up gen-cert
	./wait-for-ledger.sh
	curl http://localhost:9000/genesis > conf/genesis.txt
	docker-compose up -d

down:
	docker-compose down
	cd .docker/von-network && ./manage stop

rm: down
	docker-compose rm

