GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/main
GO_BUILD=$(shell pwd)/bin/desafio-conductor

$(DOCKER_CMD): clean
	mkdir -p $(DOCKER_BUILD)
	$(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
	rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
	heroku container:push web

test:
	go test ./...

build: test
	go build -o ${GO_BUILD}

run: build
	DB=./db/data.sqlite PORT=5000 ${GO_BUILD}
