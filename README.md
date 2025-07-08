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
2. Run `make test` to execute both the Solidity and Go tests. Behind the scenes
   this runs `forge test` and `go test ./...`.
3. Build the Go bot with `make build` which simply calls `go build ./cmd/bot`.
   Utility packages under `pkg/` keep the bot logic modular. For example,
  `ethutil.ConnectClient` provides a simple wrapper for creating Ethereum RPC
  clients. The `watcher` package offers utilities for observing chain activity.
  `BlockWatcher` logs new block headers, while `EventWatcher` can subscribe to
  contract events using a filter query.

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
pools. It takes an array of Uniswap V2 style pairs and searches for the most
profitable input amount before sequentially swapping through all pools in the
cycle. An example test demonstrates a three‑pair arbitrage executing
successfully.

Run both Solidity and Go tests with:

 ```bash
 forge test -vv
 go test ./...
 ```

Alternatively you can just run `make test` which wraps both commands.

The goal is to watch Arbitrum events for arbitrage opportunities across many
DEXes.

A sample `.env.sample` file is provided containing environment variables used by
the Go bot, such as `RPC_URL` and `PRIVATE_KEY`. Copy it to `.env` and adjust the
values as needed. `RPC_URL` defaults to the public Arbitrum RPC endpoint if left
unset.

## Docker

You can build a container image containing the bot binary:

```bash
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

\To calculate Solidity test coverage, run:

```bash
forge coverage
```

See `.env.sample` for the full list of environment variables understood by the
bot.


