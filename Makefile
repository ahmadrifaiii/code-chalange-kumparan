POST_MAIN := "$(CURDIR)/cmd/post"
BIN_POST := "$(CURDIR)/bin/post-service"
EVENT_MAIN := "$(CURDIR)/cmd/event"
BIN_EVENT := "$(CURDIR)/bin/event-service"
GET_MAIN := "$(CURDIR)/cmd/get"
BIN_GET := "$(CURDIR)/bin/get-service"


fetch:
	@go mod download

build-post:
	@go build -i -v -o $(BIN_POST) $(POST_MAIN)

build-event:
	@go build -i -v -o $(BIN_EVENT) $(EVENT_MAIN)

build-get:
	@go build -i -v -o $(BIN_GET) $(GET_MAIN)

build-post-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_POST) $(POST_MAIN)

build-event-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_EVENT) $(EVENT_MAIN)

build-get-vendor:
	@go build -mod=vendor -ldflags="-w -s" -o $(BIN_GET) $(GET_MAIN)

deploy: clean fetch build-post build-event build-get

swagger-gen:
	@swag init -g cmd/post/main.go --output pkg/swagger/docs

clean:
	@rm -rf $(CURDIR)/bin

run-test:
	@go test -v -cover ./... | grep _test.go > output.out