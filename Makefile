DC = docker-compose

run:
	$(DC) up
build:
	cd ./publisher && make docker-build && cd .. && $(DC) build && $(DC) up
