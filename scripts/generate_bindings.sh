#!/bin/bash
set -euo pipefail

if ! command -v abigen >/dev/null 2>&1; then
    echo "abigen not found, skipping bindings generation" >&2
    exit 0
fi

forge build >/dev/null

OUTDIR="cmd/bot/bindings"
mkdir -p "$OUTDIR"
CONTRACTS=(Registry ArbitrageExecutor MultiArbitrageExecutor TriangularArbitrageExecutor BatchExecutor)
for C in "${CONTRACTS[@]}"; do
    json="out/${C}.sol/${C}.json"
    abi=$(mktemp)
    bin=$(mktemp)
    jq -r '.abi' "$json" > "$abi"
    jq -r '.bytecode.object' "$json" > "$bin"
    abigen --abi "$abi" --bin "$bin" --pkg bindings --type "$C" --out "$OUTDIR/$C.go"
    rm "$abi" "$bin"
done
