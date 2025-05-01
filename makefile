SERVICE_NAME = signapse-test

-include .env
export

.PHONY: setup
setup:
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: mocks
mocks:
	mockery --all --with-expecter=true

.PHONY: fmt
fmt:
	gofmt -w -s .
	goimports -w .
	go clean ./...

.PHONY: run
run: fmt
	go run github.com/oking02/signapse-test/cmd/app
