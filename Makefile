PROJECT_NAME := duyledat197/go-template
PKG := github.com/$(PROJECT_NAME)
build:
	@go build -i -v $(PKG)/cmd/server
run:
	go run cmd/server/main.go
test:
	go test ./...
install:
	@go install ./cmd/server/.