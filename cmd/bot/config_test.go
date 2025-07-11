package main

import (
	"os"
	"testing"
)

func TestLoadConfigAndApply(t *testing.T) {
	tmp, err := os.CreateTemp(t.TempDir(), "cfg-*.toml")
	if err != nil {
		t.Fatal(err)
	}
	defer tmp.Close()
	content := `[default]
    rpc_url = "http://example"
    registry_address = "0x1"
    private_key = "abcd"
    market_cache = "db.sqlite"
    factories = "0x2"
    pairs = "0x3,0x4"
    executor_address = "0x5"
    debug = true
    `
	if _, err := tmp.WriteString(content); err != nil {
		t.Fatal(err)
	}

	cfg, err := loadConfig(tmp.Name(), "default")
	if err != nil {
		t.Fatal(err)
	}
	applyConfig(cfg)

	if os.Getenv("RPC_URL") != "http://example" ||
		os.Getenv("REGISTRY_ADDRESS") != "0x1" ||
		os.Getenv("PRIVATE_KEY") != "abcd" ||
		os.Getenv("MARKET_CACHE") != "db.sqlite" ||
		os.Getenv("FACTORIES") != "0x2" ||
		os.Getenv("PAIRS") != "0x3,0x4" ||
		os.Getenv("EXECUTOR_ADDRESS") != "0x5" ||
		os.Getenv("DEBUG") == "" {
		t.Fatal("config not applied")
	}
}
