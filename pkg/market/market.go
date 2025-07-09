package market

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Market stores discovered pool addresses.
// Market stores discovered pool and token addresses.
type Market struct {
	mu     sync.RWMutex
	pools  map[common.Address]struct{}
	tokens map[common.Address]struct{}
}

// New creates an empty Market.
func New() *Market {
	return &Market{pools: make(map[common.Address]struct{}), tokens: make(map[common.Address]struct{})}
}

// Add records a new pool address.
func (m *Market) Add(addr common.Address) {
	m.mu.Lock()
	m.pools[addr] = struct{}{}
	m.mu.Unlock()
}

// AddToken records a token address.
func (m *Market) AddToken(addr common.Address) {
	m.mu.Lock()
	m.tokens[addr] = struct{}{}
	m.mu.Unlock()
}

// Has returns true if the address is known.
func (m *Market) Has(addr common.Address) bool {
	m.mu.RLock()
	_, ok := m.pools[addr]
	m.mu.RUnlock()
	return ok
}

// HasToken returns true if the token address is known.
func (m *Market) HasToken(addr common.Address) bool {
	m.mu.RLock()
	_, ok := m.tokens[addr]
	m.mu.RUnlock()
	return ok
}

// List returns all known pool addresses.
func (m *Market) List() []common.Address {
	m.mu.RLock()
	res := make([]common.Address, 0, len(m.pools))
	for a := range m.pools {
		res = append(res, a)
	}
	m.mu.RUnlock()
	return res
}

// ListTokens returns all known token addresses.
func (m *Market) ListTokens() []common.Address {
	m.mu.RLock()
	res := make([]common.Address, 0, len(m.tokens))
	for a := range m.tokens {
		res = append(res, a)
	}
	m.mu.RUnlock()
	return res
}
