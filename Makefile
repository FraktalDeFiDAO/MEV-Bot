.RECIPEPREFIX := >
.PHONY: test solidity-test go-test build docker run run-dev deploy clean

CONTRACT ?= contracts/Registry.sol:Registry

solidity-test:
>[ -d lib/forge-std ] || forge install foundry-rs/forge-std
>forge test -vv

go-test:
>go mod download
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

deploy:
>forge create $(CONTRACT) --rpc-url $(RPC_URL) \
	--private-key $(PRIVATE_KEY) --broadcast

clean:
>rm -f bot
>rm -rf out cache
