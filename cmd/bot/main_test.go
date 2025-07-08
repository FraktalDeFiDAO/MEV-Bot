package main

import (
	"testing"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/ethutil"
)

func TestConnectClientInvalid(t *testing.T) {
	_, err := ethutil.ConnectClient(nil, "invalid://url")
	if err == nil {
		t.Fatal("expected error for invalid url")
	}
}
