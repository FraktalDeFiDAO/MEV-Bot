package bindings

import (
    "encoding/json"
    "os"
    "testing"
)

func TestRegistryArtifactContainsABI(t *testing.T) {
    data, err := os.ReadFile("../../../out/Registry.sol/Registry.json")
    if err != nil {
        t.Fatalf("failed to read artifact: %v", err)
    }
    var artifact struct {
        ABI json.RawMessage `json:"abi"`
    }
    if err := json.Unmarshal(data, &artifact); err != nil {
        t.Fatalf("failed to unmarshal json: %v", err)
    }
    if len(artifact.ABI) == 0 {
        t.Fatal("abi missing in artifact")
    }
}
