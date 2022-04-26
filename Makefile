include .env_dev

VERSION  	  	= $(shell git describe --always --tags)
NAME           	= $(shell basename $(CURDIR))

build-postgres: ##@postgres build postgres docker image
	DOCKER_BUILDKIT=1 \
	docker build \
	--progress=plain \
	-t postgres_$(NAME):$(VERSION) \
	-f ./build/postgres/Dockerfile \
	./build/postgres/
	
run-postgres: build-postgres  ##@postgres run postgres on docker
	DOCKER_BUILDKIT=1 \
	docker run --rm -d \
	-v $(HOME)/admin-api/pgdata:/var/lib/postgresql/mydata \
	-p 5434:5432 \
	--name postgres_$(NAME):$(VERSION) \
	postgres_$(NAME):$(VERSION)
	