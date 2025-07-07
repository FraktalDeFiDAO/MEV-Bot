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
