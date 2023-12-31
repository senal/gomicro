SHELL=cmd.exe
BROKER_BINARY=brokerApp
AUTHENTICATION_BINARY=authenticationApp

## up: starts all containers in the background without forcing build
up: 
	@echo Starting Docker images...
	docker-compose up -d
	@echo Docker images started!

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_authentication 
	@echo Stopping docker images (if running...)
	docker-compose down
	@echo Building (when required) and starting docker images...
	docker-compose up --build -d
	@echo Docker images built and started!

## down: stop docker compose
down:
	@echo Stopping docker compose...
	docker-compose down
	@echo Done!

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo Building broker binary...
	chdir broker && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${BROKER_BINARY} ./cmd/api
	@echo Done!

## build_authentication: builds the authentication binary as a linux executable
build_authentication:
	@echo Building authentication binary...
	chdir authentication && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0&& go build -o ${AUTHENTICATION_BINARY} ./cmd/api
	@echo Done!
