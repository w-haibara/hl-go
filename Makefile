hl: cmd/hl/main.go go.mod
	go fmt ./...
	go mod tidy
	go build -o hl ./cmd/...

.PHONY: run
run:
	go fmt ./...
	go mod tidy
	go run ./cmd/...

.PHONY: init
init:
	go mod init hl
	go mod tidy

.PHONY: test
test:
	go fmt ./...
	go test -v
