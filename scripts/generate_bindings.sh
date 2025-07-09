#!/bin/bash
set -euo pipefail
OUTDIR="cmd/bot/bindings"
mkdir -p "$OUTDIR"
CONTRACTS=(Registry ArbitrageExecutor MultiArbitrageExecutor TriangularArbitrageExecutor BatchExecutor)
for C in "${CONTRACTS[@]}"; do
    abigen --sol "contracts/$C.sol" --pkg bindings --out "$OUTDIR/$C.go"
done
