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
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

// Client wraps calls to the on-chain Registry contract.
type Client struct {
	c    contract
	auth *bind.TransactOpts
}

const abiJSON = `[
{"inputs":[{"internalType":"address","name":"token","type":"address"},{"internalType":"uint8","name":"decimals","type":"uint8"}],"name":"addToken","outputs":[],"stateMutability":"nonpayable","type":"function"},
{"inputs":[{"internalType":"address","name":"pool","type":"address"},{"internalType":"address","name":"token0","type":"address"},{"internalType":"address","name":"token1","type":"address"},{"internalType":"uint256","name":"exchangeId","type":"uint256"}],"name":"addPool","outputs":[],"stateMutability":"nonpayable","type":"function"}
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
	return &Client{c: bc, auth: auth}, nil
}

// AddPool submits a transaction to record a pool in the registry.
func (r *Client) AddPool(pool, token0, token1 common.Address, exchangeID uint64) (*types.Transaction, error) {
	return r.c.Transact(r.auth, "addPool", pool, token0, token1, big.NewInt(int64(exchangeID)))
}

// AddToken submits a transaction to record a token in the registry.
func (r *Client) AddToken(token common.Address, decimals uint8) (*types.Transaction, error) {
	return r.c.Transact(r.auth, "addToken", token, decimals)
}
