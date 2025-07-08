.PHONY: test solidity-test go-test build

solidity-test:
	forge test -vv

go-test:
	go test ./...

test: solidity-test go-test

build:
	go build ./cmd/bot
