.RECIPEPREFIX := >
.PHONY: test solidity-test go-test build docker run run-dev debug registry-cli deploy clean generate-bindings web-install web-build web-dev

CONTRACT ?= contracts/Registry.sol:Registry

solidity-test:
>[ -d lib/forge-std ] || forge install foundry-rs/forge-std
>forge test -vv

go-test:
>go mod download
>go test ./...

test: solidity-test go-test

build: generate-bindings
>go build ./cmd/bot

docker:
>docker build -t mev-bot .

run: build
>./bot

run-dev: generate-bindings
>go run ./cmd/bot

debug: generate-bindings
>dlv debug ./cmd/bot -- $(ARGS)

registry-cli:
>go run ./cmd/registry $(ARGS)

deploy:
>forge create $(CONTRACT) --rpc-url $(RPC_URL) --private-key $(PRIVATE_KEY)

generate-bindings:
>scripts/generate_bindings.sh

clean:
>rm -f bot
>rm -rf out cache

web-install:
>npm install --prefix web

web-build:
>npm run build --prefix web

web-dev:
>npm run dev --prefix web
