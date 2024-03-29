BINARY_NAME=handler

.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

audit:
	go mod verify
	go vet ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest cmd/

clean:
	go clean

test:
	go test -v cmd/main.go

deps:
	go get go.uber.org/zap

build:
	GOOS=linux GOARCH=amd64 go build -o ${BINARY_NAME} cmd/main.go

run:
	func start

all:
	tidy build