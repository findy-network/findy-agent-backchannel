von-up:
	mkdir -p .docker
	-git clone https://github.com/bcgov/von-network .docker/von-network
	cd .docker/von-network && ./manage build && ./manage start

aath-up: rm up fth-up
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness
#	cd .docker/aries-agent-test-harness && \
#		./manage build -a acapy && \
#		./manage run -d acapy -t @T001-RFC0160
	cd .docker/aries-agent-test-harness && \
		mkdir -p aries-backchannels/findy && \
		touch aries-backchannels/findy/Dockerfile.findy && \
		./manage run -d findy -b acapy -t @T001-RFC0160

aath-only: fth-up
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness
	cd .docker/aries-agent-test-harness && \
		mkdir -p aries-backchannels/findy && \
		touch aries-backchannels/findy/Dockerfile.findy && \
		./manage run -d acapy -b findy -t @T001-RFC0160

fth-up:
# build and run docker images with correct name
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
	sleep 15
	curl http://localhost:9000/genesis > conf/genesis.txt
	docker-compose up -d

down:
	docker-compose down
	cd .docker/von-network && ./manage stop

rm: down
	docker-compose rm

