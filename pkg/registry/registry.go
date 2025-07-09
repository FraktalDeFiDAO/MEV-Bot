package registry

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type contract interface {
	Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

// Client wraps calls to the on-chain Registry contract.
type Client struct {
	c       contract
	rpc     *ethclient.Client
	auth    *bind.TransactOpts
	callCtx *bind.CallOpts
}

type PoolInfo struct {
	Token0     common.Address
	Token1     common.Address
	ExchangeID *big.Int
	Enabled    bool
}

const abiJSON = `[
{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint8","name":"decimals","type":"uint8"}],"name":"addToken","outputs":[],"stateMutability":"nonpayable","type":"function"},
{"inputs":[{"internalType":"address","name":"pool","type":"address"},{"internalType":"address","name":"token0","type":"address"},{"internalType":"address","name":"token1","type":"address"},{"internalType":"uint256","name":"exchangeId","type":"uint256"}],"name":"addPool","outputs":[],"stateMutability":"nonpayable","type":"function"},
{"inputs":[],"name":"getTokens","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
{"inputs":[],"name":"getPools","outputs":[{"internalType":"address[]","name":"","type":"address[]"}],"stateMutability":"view","type":"function"},
{"inputs":[{"internalType":"address","name":"pool","type":"address"}],"name":"getPool","outputs":[{"components":[{"internalType":"address","name":"token0","type":"address"},{"internalType":"address","name":"token1","type":"address"},{"internalType":"uint256","name":"exchangeId","type":"uint256"},{"internalType":"bool","name":"enabled","type":"bool"}],"internalType":"struct PoolLib.PoolInfo","name":"","type":"tuple"}],"stateMutability":"view","type":"function"}
]`

// New creates a client for the registry at addr using the given ethclient and private key.
func New(ctx context.Context, addr common.Address, rpc *ethclient.Client, key *ecdsa.PrivateKey) (*Client, error) {
	parsed, err := abi.JSON(strings.NewReader(abiJSON))
	if err != nil {
		return nil, err
	}
	chainID, err := rpc.ChainID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		return nil, err
	}
	bc := bind.NewBoundContract(addr, parsed, rpc, rpc, rpc)
	return &Client{c: bc, rpc: rpc, auth: auth, callCtx: &bind.CallOpts{Context: ctx}}, nil
}

// AddPool submits a transaction to record a pool in the registry.
func (r *Client) AddPool(pool, token0, token1 common.Address, exchangeID uint64) (*types.Transaction, error) {
	return r.c.Transact(r.auth, "addPool", pool, token0, token1, big.NewInt(int64(exchangeID)))
}

// AddToken submits a transaction to record a token in the registry.
func (r *Client) AddToken(token common.Address, decimals uint8) (*types.Transaction, error) {
	return r.c.Transact(r.auth, "addToken", token, decimals)
}

// WaitMined blocks until the transaction is mined and returns the receipt.
func (r *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, r.rpc, tx)
}

// Tokens returns the list of registered token addresses.
func (r *Client) Tokens(ctx context.Context) ([]common.Address, error) {
	var raw []interface{}
	if err := r.c.Call(&bind.CallOpts{Context: ctx}, &raw, "getTokens"); err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, nil
	}
	return raw[0].([]common.Address), nil
}

// Pools returns the list of registered pool addresses.
func (r *Client) Pools(ctx context.Context) ([]common.Address, error) {
	var raw []interface{}
	if err := r.c.Call(&bind.CallOpts{Context: ctx}, &raw, "getPools"); err != nil {
		return nil, err
	}
	if len(raw) == 0 {
		return nil, nil
	}
	return raw[0].([]common.Address), nil
}

// PoolInfo returns metadata for a given pool.
func (r *Client) PoolInfo(ctx context.Context, pool common.Address) (PoolInfo, error) {
	var raw []interface{}
	if err := r.c.Call(&bind.CallOpts{Context: ctx}, &raw, "getPool", pool); err != nil {
		return PoolInfo{}, err
	}
	return PoolInfo{
		Token0:     raw[0].(common.Address),
		Token1:     raw[1].(common.Address),
		ExchangeID: raw[2].(*big.Int),
		Enabled:    raw[3].(bool),
	}, nil
}
