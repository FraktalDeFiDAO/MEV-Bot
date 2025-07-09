# MEV-Bot

This project explores building a MEV bot targeting decentralized exchanges on Arbitrum.

- **Smart contracts** live in the `./contracts` directory. They are compiled and tested using [Foundry](https://github.com/foundry-rs/foundry).
- **Foundry configuration** is contained in `foundry.toml` at the repository root. Run tests with `forge test`.
- **Bot implementation** will be written in Go. The entry point is under `cmd/bot` and connects to an Arbitrum node to watch transactions and events.

## Getting started

1. Install [Foundry](https://book.getfoundry.sh/getting-started/installation) using:
   ```bash
   curl -L https://foundry.paradigm.xyz | bash
   source ~/.bashrc && foundryup
   ```
   After cloning the repo, run `forge install foundry-rs/forge-std` to fetch
   the testing libraries such as `forge-std`.
   The `lib` directory is excluded from version control, so you may need to
   run this command whenever you clone a fresh repository.
2. Download Go dependencies with `go mod download` so the Go tests can compile.
3. Run `make test` to execute both the Solidity and Go tests. Behind the scenes
   this runs `forge test` and `go test ./...`.
4. Build the Go bot with `make build` which simply calls `go build ./cmd/bot`.
   Utility packages under `pkg/` keep the bot logic modular. For example,
  `ethutil.ConnectClient` provides a simple wrapper for creating Ethereum RPC
  clients. The `watcher` package offers utilities for observing chain activity.
  `BlockWatcher` logs new block headers only when `DEBUG=1` is set to reduce
  noise, while `EventWatcher` subscribes only to
  profitable trade events and Uniswap V2 `Sync` events. When a
  `TradeExecuted` event is observed the bot prints the input and profit so
  profitable trades are logged in real time. `Sync` events are decoded to
  log pool price updates, allowing the bot to track market movements without
  processing every log on the chain.
5. Start the bot with `make run` (or `make run-dev` to run using `go run`).
   The bot automatically loads environment variables from a `.env` file if
   present using `godotenv`. Set `DEBUG=1` to enable more verbose logging with
   file and line numbers. **Use a WebSocket RPC URL** (e.g. `wss://...`) so the
   block and event watchers can subscribe to notifications. HTTP endpoints will
   only log an error and no events will be seen. The bot will look for a `PAIRS`
   environment variable specifying pairs to monitor for arbitrage, formatted as
   `"addr1,addr2;addr3,addr4"`.
   Set `FACTORIES` to a comma separated list of factory addresses to
   automatically capture `PairCreated` and `PoolCreated` events and grow the
   set of pools scanned for opportunities.
6. Deploy contracts using `make deploy`. By default this deploys the
   `Registry` contract with `forge create`.  Pass `CONTRACT=path:Name` to
   deploy a different contract.  `RPC_URL` and `PRIVATE_KEY` must be set in
   the environment.
7. Generate Go contract bindings with `make generate-bindings`. This runs
   `scripts/generate_bindings.sh` which uses `abigen` to create Go packages
   under `cmd/bot/bindings`. If `abigen` isn't installed, you can add it to
   your `PATH` with:

   ```bash
   go install github.com/ethereum/go-ethereum/cmd/abigen@latest



The repo now includes a `Registry` contract that stores token, exchange and pool metadata using library based diamond storage. It forms the on-chain
configuration for the bot and demonstrates how components remain modular.
The registry exposes helper getters like `getTokenCount`, `getExchangeCount` and
`getPoolCount` so the bot can query how many entries exist without fetching all
data.

`BatchExecutor` provides a simple multicall style contract that allows the bot
to execute a series of encoded calls within a single transaction. This makes it
easy to chain swaps across multiple pools when seeking arbitrage opportunities.

`ArbitrageCalculator` implements a small library that estimates the most
profitable trade size between two constant product pools.  It helps the bot
determine how much input to supply when performing cross-exchange arbitrage.

`ArbitrageExecutor` is a simple contract that uses this calculator to execute
an arbitrage across two constant product pools. The executor transfers tokens
between the pairs and performs the swaps, returning any profit to the caller.

`TriangularArbitrageCalculator` and `TriangularArbitrageExecutor` extend this
idea to three pools. The calculator searches for the most profitable input
amount when swapping through a cycle of three constant product pairs. The
executor performs the swaps atomically, enabling triangular arbitrage across
multiple DEXes.

`MultiArbitrageExecutor` further generalizes this to an arbitrary number of
pools. It now accepts per‑pool fee settings so different exchanges like Uniswap
V3 or Algebra can be mixed within a single cycle. Provide parallel arrays of fee
numerators and denominators when calling `execute` to support directional fees.
An example test demonstrates a three‑pair arbitrage executing successfully.

Run both Solidity and Go tests with:

 ```bash
 forge test -vv
 go test ./...
 ```

Alternatively you can just run `make test` which wraps both commands.

The goal is to watch Arbitrum events for arbitrage opportunities across many
DEXes.

`Market` now records both discovered pools and the tokens they trade. Pool and
token addresses are collected from factory events in real time and can be
listed via the bot's API in future iterations. A simple `NonceManager` utility
helps avoid transaction conflicts when submitting many arbitrage attempts.

A lightweight Vue 3 front end scaffold lives in `web/`. It uses Vite,
Tailwind, Pinia, Viem, and Vue Router to visualize saved tokens and pools.
Install dependencies with `make web-install` and run `make web-dev` to launch
the UI.  Build the production bundle with `make web-build`.

A sample `.env.sample` file is provided containing environment variables used by
the Go bot, such as `RPC_URL` and `PRIVATE_KEY`. Copy it to `.env` and adjust the
values as needed.  `RPC_URL` defaults to the public Arbitrum RPC endpoint if left
unset. `MARKET_CACHE` controls where discovered pools are cached locally and
`REGISTRY_ADDRESS` specifies an on-chain registry contract to persist newly
discovered pools and tokens.

## Docker

You can build a container image containing the bot binary:

```bash
# or simply `make docker`
docker build -t mev-bot .
```

Run the container with your environment variables, for example:

```bash
docker run --rm -e RPC_URL=https://arb1.arbitrum.io/rpc mev-bot
```

To check Go test coverage you can run:

```bash
go test ./... -cover
```

To calculate Solidity test coverage, run:

```bash
forge coverage
```

See `.env.sample` for the full list of environment variables understood by the
bot.


