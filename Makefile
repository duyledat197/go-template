PROJECT_NAME := "stamp-server"
PKG := "github.com/$(PROJECT_NAME)"
build:
	@go build -i -v $(PKG)/cmd/server