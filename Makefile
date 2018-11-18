DC = docker-compose

run:
	$(DC) up
build: build-pub build-sub
	 $(DC) build && $(DC) up
build-pub: ./publisher/*.go
	cd ./publisher && make docker-build && cd ..
build-sub: ./subscriber/*.go
	cd ./subscriber && make docker-build && cd ..
