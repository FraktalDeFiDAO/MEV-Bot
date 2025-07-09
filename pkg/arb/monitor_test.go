package arb

import (
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

type logCatcher struct{ msgs int }

func (l *logCatcher) Write(p []byte) (n int, err error) {
	l.msgs++
	return len(p), nil
}

func TestMonitorDetects(t *testing.T) {
	addrA := common.HexToAddress("0x1")
	addrB := common.HexToAddress("0x2")
	m := NewMonitor([][2]common.Address{{addrA, addrB}}, 500, 1)
	logger := &logCatcher{}
	log.SetOutput(logger)
	defer log.SetOutput(nil)

	m.Update(addrA, big.NewInt(1000), big.NewInt(1000))
	m.Update(addrB, big.NewInt(1200), big.NewInt(800))
	if logger.msgs == 0 {
		t.Fatalf("expected log message")
	}
}
