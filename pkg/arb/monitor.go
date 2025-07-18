package arb

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Pool struct {
	Res0 *big.Int
	Res1 *big.Int
}

type Monitor struct {
	pairs [][2]common.Address
	pools map[common.Address]*Pool
	maxIn int64
	step  int64
}

func NewMonitor(pairs [][2]common.Address, maxIn, step int64) *Monitor {
	return &Monitor{pairs: pairs, pools: make(map[common.Address]*Pool), maxIn: maxIn, step: step}
}

func (m *Monitor) Update(addr common.Address, r0, r1 *big.Int) {
	m.pools[addr] = &Pool{new(big.Int).Set(r0), new(big.Int).Set(r1)}
	m.check()
}

// AddPair registers a new pair to monitor.
func (m *Monitor) AddPair(a, b common.Address) {
	for _, p := range m.pairs {
		if (p[0] == a && p[1] == b) || (p[0] == b && p[1] == a) {
			return
		}
	}
	m.pairs = append(m.pairs, [2]common.Address{a, b})
	m.check()
}

func (m *Monitor) check() {
	for _, p := range m.pairs {
		a := m.pools[p[0]]
		b := m.pools[p[1]]
		if a == nil || b == nil {
			continue
		}
		bestIn, profit := FindBestInput(a.Res0, a.Res1, b.Res0, b.Res1, m.maxIn, m.step)
		if profit.Cmp(big.NewInt(0)) > 0 {
			log.Printf("arbitrage opportunity %s -> %s amount=%s profit=%s", p[0].Hex(), p[1].Hex(), bestIn, profit)
		}
	}
}
