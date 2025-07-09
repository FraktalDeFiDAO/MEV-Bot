package registry

import (
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type stubContract struct {
	method string
	params []interface{}
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
