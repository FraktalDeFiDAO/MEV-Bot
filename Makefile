.RECIPEPREFIX := >
.PHONY: test solidity-test go-test build docker run run-dev

solidity-test:
>forge test -vv

go-test:
>go test ./...

test: solidity-test go-test

build:
>go build ./cmd/bot

docker:
>docker build -t mev-bot .

run: build
>./bot

run-dev:
>go run ./cmd/bot
