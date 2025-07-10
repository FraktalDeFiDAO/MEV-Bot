package registry

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type stubContract struct {
	method  string
	params  []interface{}
	callRes interface{}
}

func (s *stubContract) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	s.method = method
	s.params = params
	switch v := s.callRes.(type) {
	case []common.Address:
		*result = append(*result, v)
	case PoolInfo:
		tup := struct {
			Token0     common.Address `json:"token0"`
			Token1     common.Address `json:"token1"`
			ExchangeId *big.Int       `json:"exchangeId"`
			Enabled    bool           `json:"enabled"`
		}{v.Token0, v.Token1, v.ExchangeID, v.Enabled}
		*result = append(*result, tup)
	case struct {
		Token0     common.Address `json:"token0"`
		Token1     common.Address `json:"token1"`
		ExchangeId *big.Int       `json:"exchangeId"`
		Enabled    bool           `json:"enabled"`
	}:
		*result = append(*result, struct {
			Token0     common.Address `json:"token0"`
			Token1     common.Address `json:"token1"`
			ExchangeId *big.Int       `json:"exchangeId"`
			Enabled    bool           `json:"enabled"`
		}{v.Token0, v.Token1, v.ExchangeId, v.Enabled})
	}
	return nil
}

func (s *stubContract) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	s.method = method
	s.params = params
	return new(types.Transaction), nil
}

func TestAddToken(t *testing.T) {
	st := &stubContract{}
	c := &Client{c: st, auth: &bind.TransactOpts{}}
	if _, err := c.AddToken(common.HexToAddress("0x1"), 18); err != nil {
		t.Fatal(err)
	}
	if st.method != "addToken" {
		t.Fatalf("wrong method %s", st.method)
	}
	if len(st.params) != 2 {
		t.Fatalf("wrong params %v", st.params)
	}
}

func TestAddPool(t *testing.T) {
	st := &stubContract{}
	c := &Client{c: st, auth: &bind.TransactOpts{}}
	if _, err := c.AddPool(common.HexToAddress("0x1"), common.HexToAddress("0x2"), common.HexToAddress("0x3"), 1); err != nil {
		t.Fatal(err)
	}
	if st.method != "addPool" {
		t.Fatalf("wrong method %s", st.method)
	}
	if len(st.params) != 4 {
		t.Fatalf("wrong params %v", st.params)
	}
}

func TestTokens(t *testing.T) {
	st := &stubContract{callRes: []common.Address{common.HexToAddress("0x1")}}
	c := &Client{c: st, callCtx: &bind.CallOpts{}}
	toks, err := c.Tokens(context.Background())
	if err != nil || len(toks) != 1 || toks[0] != common.HexToAddress("0x1") {
		t.Fatalf("unexpected tokens %v %v", toks, err)
	}
	if st.method != "getTokens" {
		t.Fatalf("wrong method %s", st.method)
	}
}

func TestPoolInfo(t *testing.T) {
	out := struct {
		Token0     common.Address `json:"token0"`
		Token1     common.Address `json:"token1"`
		ExchangeId *big.Int       `json:"exchangeId"`
		Enabled    bool           `json:"enabled"`
	}{common.HexToAddress("0x1"), common.HexToAddress("0x2"), big.NewInt(1), true}
	st := &stubContract{callRes: out}
	c := &Client{c: st, callCtx: &bind.CallOpts{}}
	info, err := c.PoolInfo(context.Background(), common.HexToAddress("0xabc"))
	if err != nil || info.Token0 != out.Token0 || info.Token1 != out.Token1 {
		t.Fatalf("unexpected info %+v %v", info, err)
	}
	if st.method != "getPool" {
		t.Fatalf("wrong method %s", st.method)
	}
}
