package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func rpcTestServer() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"jsonrpc":"2.0","id":1,"result":"0x1"}`))
	}))
}

func TestConnectMissingEnv(t *testing.T) {
	os.Unsetenv("REGISTRY_ADDRESS")
	os.Unsetenv("PRIVATE_KEY")
	if _, _, err := connect(context.Background()); err == nil {
		t.Fatal("expected error when env vars missing")
	}
}

func TestConnectSuccess(t *testing.T) {
	srv := rpcTestServer()
	defer srv.Close()

	os.Setenv("RPC_URL", srv.URL)
	os.Setenv("REGISTRY_ADDRESS", "0x0000000000000000000000000000000000000001")
	os.Setenv("PRIVATE_KEY", "4f3edf983ac636a65a842ce7c78d9aa706d3b113b37e7e5ccef5d266c131c7da")
	defer os.Unsetenv("RPC_URL")
	defer os.Unsetenv("REGISTRY_ADDRESS")
	defer os.Unsetenv("PRIVATE_KEY")

	if _, _, err := connect(context.Background()); err != nil {
		t.Fatalf("connect failed: %v", err)
	}
}

func TestConnectInvalidKey(t *testing.T) {
	srv := rpcTestServer()
	defer srv.Close()

	os.Setenv("RPC_URL", srv.URL)
	os.Setenv("REGISTRY_ADDRESS", "0x0000000000000000000000000000000000000001")
	os.Setenv("PRIVATE_KEY", "badkey")
	defer os.Unsetenv("RPC_URL")
	defer os.Unsetenv("REGISTRY_ADDRESS")
	defer os.Unsetenv("PRIVATE_KEY")

	if _, _, err := connect(context.Background()); err == nil {
		t.Fatal("expected error for invalid key")
	}
}
