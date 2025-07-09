#!/bin/bash
set -euo pipefail

if ! command -v abigen >/dev/null 2>&1; then
    echo "abigen not found; install it with 'go install github.com/ethereum/go-ethereum/cmd/abigen@latest'" >&2
    exit 1
fi
OUTDIR="cmd/bot/bindings"
mkdir -p "$OUTDIR"
CONTRACTS=(Registry ArbitrageExecutor MultiArbitrageExecutor TriangularArbitrageExecutor BatchExecutor)
for C in "${CONTRACTS[@]}"; do
    abigen --sol "contracts/$C.sol" --pkg bindings --out "$OUTDIR/$C.go"
done
