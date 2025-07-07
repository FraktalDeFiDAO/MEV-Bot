# MEV-Bot

This project explores building a MEV bot targeting decentralized exchanges on Arbitrum.

- **Smart contracts** live in the `./contracts` directory. They are compiled and tested using [Foundry](https://github.com/foundry-rs/foundry).
- **Foundry configuration** is contained in `foundry.toml` at the repository root. Run tests with `forge test`.
- **Bot implementation** will be written in Go. The entry point is under `cmd/bot` and connects to an Arbitrum node to watch transactions and events.

## Getting started

1. Install [Foundry](https://book.getfoundry.sh/getting-started/installation).
2. Run `forge test` to compile and test the contracts.
3. Build the Go bot with `go build ./cmd/bot`.

This is a minimal MVP setup aiming to monitor Arbitrum transactions and eventually the sequencer for opportunities across multiple DEXes.

The repo now includes a `Registry` contract that stores token, exchange and pool metadata using library based diamond storage. It forms the on-chain
configuration for the bot and demonstrates how components remain modular.

Run both Solidity and Go tests with:

```bash
forge test -vv
go test ./...
```

The goal is to watch Arbitrum events for arbitrage opportunities across many
DEXes.
>>>>>>> s9mh77-codex/create-foundry.toml-for-solidity-testing
