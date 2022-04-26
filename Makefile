include .env_dev

NAME = $(shell basename $(CURDIR))

build-postgres: ##@postgres build postgres docker image
	DOCKER_BUILDKIT=1 \
	docker build \
	--progress=plain \
	-t postgres_$(NAME) \
	-f ./postgres/Dockerfile \
	./postgres/
	
run-postgres: build-postgres  ##@postgres run postgres on docker
	DOCKER_BUILDKIT=1 \
	docker run --rm \
	-v $(HOME)/hotelsapi/pgdata:/var/lib/postgresql/data \
	-p 5432:5432 \
	postgres_$(NAME)
	
run:
	POSTGRES_URL=$(POSTGRES_URL) \
	go run cmd/main.go