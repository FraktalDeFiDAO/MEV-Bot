package market

import (
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// Market stores discovered pool addresses.
// Market stores discovered pool and token addresses.
// PoolEntry holds metadata for a discovered pool.
type PoolEntry struct {
	Address common.Address
	Token0  common.Address
	Token1  common.Address
}

// Market stores discovered pools and tokens.
type Market struct {
	mu     sync.RWMutex
	pools  map[common.Address]PoolEntry
	tokens map[common.Address]struct{}
}

// New creates an empty Market.
func New() *Market {
	return &Market{pools: make(map[common.Address]PoolEntry), tokens: make(map[common.Address]struct{})}
}

// Add records a new pool address without token metadata.
func (m *Market) Add(addr common.Address) {
	m.mu.Lock()
	if _, ok := m.pools[addr]; !ok {
		m.pools[addr] = PoolEntry{Address: addr}
	}
	m.mu.Unlock()
}

// AddPool records a pool along with its tokens.
func (m *Market) AddPool(addr, token0, token1 common.Address) {
	m.mu.Lock()
	m.pools[addr] = PoolEntry{Address: addr, Token0: token0, Token1: token1}
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

// GetPool returns the tokens for a given pool.
func (m *Market) GetPool(addr common.Address) (PoolEntry, bool) {
	m.mu.RLock()
	p, ok := m.pools[addr]
	m.mu.RUnlock()
	return p, ok
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

// ListPools returns all known pools with their tokens.
func (m *Market) ListPools() []PoolEntry {
	m.mu.RLock()
	res := make([]PoolEntry, 0, len(m.pools))
	for _, p := range m.pools {
		res = append(res, p)
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

// PoolsForTokens returns pools that trade the given token pair.
func (m *Market) PoolsForTokens(t0, t1 common.Address) []common.Address {
	m.mu.RLock()
	defer m.mu.RUnlock()
	var res []common.Address
	for addr, p := range m.pools {
		if (p.Token0 == t0 && p.Token1 == t1) || (p.Token0 == t1 && p.Token1 == t0) {
			res = append(res, addr)
		}
	}
	return res
}
