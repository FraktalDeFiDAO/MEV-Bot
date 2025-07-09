package market

import (
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

// Persistent wraps Market and stores data on disk.
type Persistent struct {
	*Market
	path string
}

// LoadFromFile loads market data from a JSON file if it exists.
func LoadFromFile(path string) *Persistent {
	pm := &Persistent{Market: New(), path: path}
	data, err := os.ReadFile(path)
	if err != nil {
		return pm
	}
	var s struct {
		Pools  []string `json:"pools"`
		Tokens []string `json:"tokens"`
	}
	if json.Unmarshal(data, &s) == nil {
		for _, p := range s.Pools {
			pm.Market.Add(common.HexToAddress(p))
		}
		for _, t := range s.Tokens {
			pm.Market.AddToken(common.HexToAddress(t))
		}
	}
	return pm
}

func (p *Persistent) save() {
	var s struct {
		Pools  []string `json:"pools"`
		Tokens []string `json:"tokens"`
	}
	for _, a := range p.Market.List() {
		s.Pools = append(s.Pools, a.Hex())
	}
	for _, t := range p.Market.ListTokens() {
		s.Tokens = append(s.Tokens, t.Hex())
	}
	if data, err := json.Marshal(s); err == nil {
		_ = os.WriteFile(p.path, data, 0644)
	}
}

// Add records a pool and saves to disk if not already present.
func (p *Persistent) Add(addr common.Address) {
	if !p.Market.Has(addr) {
		p.Market.Add(addr)
		p.save()
	}
}

// AddToken records a token and saves to disk if not present.
func (p *Persistent) AddToken(addr common.Address) {
	if !p.Market.HasToken(addr) {
		p.Market.AddToken(addr)
		p.save()
	}
}
