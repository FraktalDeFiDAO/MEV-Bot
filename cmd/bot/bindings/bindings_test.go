package bindings

import "testing"

func TestRegistryABILoad(t *testing.T) {
	if _, err := RegistryMetaData.GetAbi(); err != nil {
		t.Fatalf("failed to parse registry ABI: %v", err)
	}
}
