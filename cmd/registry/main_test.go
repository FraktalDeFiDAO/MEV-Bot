package main

import (
	"context"
	"math/big"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/FraktalDeFiDAO/MEV-Bot/pkg/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

type stubClient struct {
	tokens      []common.Address
	pools       []common.Address
	poolInfo    registry.PoolInfo
	addedTokens []common.Address
	addedPools  []struct {
		pool common.Address
		t0   common.Address
		t1   common.Address
		id   uint64
	}
}

func (s *stubClient) AddPool(a, b, c common.Address, id uint64) (*types.Transaction, error) {
	s.addedPools = append(s.addedPools, struct {
		pool common.Address
		t0   common.Address
		t1   common.Address
		id   uint64
	}{a, b, c, id})
	return types.NewTx(&types.LegacyTx{Nonce: 0}), nil
}

func (s *stubClient) AddToken(a common.Address, d uint8) (*types.Transaction, error) {
	s.addedTokens = append(s.addedTokens, a)
	return types.NewTx(&types.LegacyTx{Nonce: 0}), nil
}

func (s *stubClient) WaitMined(context.Context, *types.Transaction) (*types.Receipt, error) {
	return &types.Receipt{Status: types.ReceiptStatusSuccessful}, nil
}

func (s *stubClient) Tokens(context.Context) ([]common.Address, error) { return s.tokens, nil }
func (s *stubClient) Pools(context.Context) ([]common.Address, error)  { return s.pools, nil }
func (s *stubClient) PoolInfo(context.Context, common.Address) (registry.PoolInfo, error) {
	return s.poolInfo, nil
}

func TestHandleTokens(t *testing.T) {
	c := &stubClient{tokens: []common.Address{common.HexToAddress("0x1")}}
	if err := handle(context.Background(), c, nil, []string{"tokens"}); err != nil {
		t.Fatal(err)
	}
}

func TestHandlePools(t *testing.T) {
	c := &stubClient{
		pools:    []common.Address{common.HexToAddress("0x1")},
		poolInfo: registry.PoolInfo{Token0: common.HexToAddress("0x2"), Token1: common.HexToAddress("0x3"), ExchangeID: big.NewInt(1)},
	}
	if err := handle(context.Background(), c, nil, []string{"pools"}); err != nil {
		t.Fatal(err)
	}
}

func TestHandleAddToken(t *testing.T) {
	c := &stubClient{}
	if err := handle(context.Background(), c, nil, []string{"add-token", "0x1"}); err != nil {
		t.Fatal(err)
	}
	if len(c.addedTokens) != 1 || c.addedTokens[0] != common.HexToAddress("0x1") {
		t.Fatalf("token not added: %v", c.addedTokens)
	}
}

func TestHandleAddPool(t *testing.T) {
	c := &stubClient{}
	pool := common.HexToAddress("0x111")
	t0 := common.HexToAddress("0x222")
	t1 := common.HexToAddress("0x333")
	if err := handle(context.Background(), c, nil, []string{"add-pool", pool.Hex(), t0.Hex(), t1.Hex(), "5"}); err != nil {
		t.Fatal(err)
	}
	if len(c.addedPools) != 1 {
		t.Fatalf("pool not added: %v", c.addedPools)
	}
	ap := c.addedPools[0]
	if ap.pool != pool || ap.t0 != t0 || ap.t1 != t1 || ap.id != 5 {
		t.Fatalf("unexpected add pool: %+v", ap)
	}
}
