von-up:
	mkdir -p .docker
	-git clone https://github.com/bcgov/von-network .docker/von-network
	cd .docker/von-network && ./manage build && ./manage start

aath-up: up fth-up
	-git clone https://github.com/hyperledger/aries-agent-test-harness .docker/aries-agent-test-harness
#	cd .docker/aries-agent-test-harness && \
#		./manage build -a acapy && \
#		./manage run -d acapy -t @T001-RFC0160
	cd .docker/aries-agent-test-harness && \
		./manage run -d findy -t @T001-RFC0160

fth-up:
# build and run docker images with correct name

up: von-up
	curl http://localhost:9000/genesis > conf/genesis.txt
	docker-compose up

down:
	docker-compose down
	cd .docker/von-network && ./manage stop

rm: down
	docker-compose rm

