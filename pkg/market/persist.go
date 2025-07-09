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
		Pools []struct {
			Address string `json:"address"`
			Token0  string `json:"token0"`
			Token1  string `json:"token1"`
		} `json:"pools"`
		Tokens []string `json:"tokens"`
	}
	if json.Unmarshal(data, &s) == nil {
		for _, pinfo := range s.Pools {
			addr := common.HexToAddress(pinfo.Address)
			t0 := common.HexToAddress(pinfo.Token0)
			t1 := common.HexToAddress(pinfo.Token1)
			if (t0 == common.Address{} && t1 == common.Address{}) {
				pm.Market.Add(addr)
			} else {
				pm.Market.AddPool(addr, t0, t1)
			}
		}
		for _, t := range s.Tokens {
			pm.Market.AddToken(common.HexToAddress(t))
		}
	}
	return pm
}

func (p *Persistent) save() {
	var s struct {
		Pools []struct {
			Address string `json:"address"`
			Token0  string `json:"token0"`
			Token1  string `json:"token1"`
		} `json:"pools"`
		Tokens []string `json:"tokens"`
	}
	for _, pe := range p.Market.ListPools() {
		s.Pools = append(s.Pools, struct {
			Address string `json:"address"`
			Token0  string `json:"token0"`
			Token1  string `json:"token1"`
		}{pe.Address.Hex(), pe.Token0.Hex(), pe.Token1.Hex()})
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

// AddPool records a pool with token metadata and saves to disk.
func (p *Persistent) AddPool(addr, token0, token1 common.Address) {
	if !p.Market.Has(addr) {
		p.Market.AddPool(addr, token0, token1)
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
