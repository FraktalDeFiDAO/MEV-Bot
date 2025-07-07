package main

import "testing"

func TestConnectClientInvalid(t *testing.T) {
    _, err := ConnectClient(nil, "invalid://url")
    if err == nil {
        t.Fatal("expected error for invalid url")
    }
}
