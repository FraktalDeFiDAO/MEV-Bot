package main

import (
	"context"
	"testing"
)

func TestRunInvalidURL(t *testing.T) {
	if err := run(context.Background(), "http://127.0.0.1:0"); err == nil {
		t.Fatal("expected error from invalid RPC URL")
	}
}
