package market

import (
	"database/sql"
	"log"

	"github.com/ethereum/go-ethereum/common"
	_ "github.com/mattn/go-sqlite3"
)

// Persistent wraps Market and stores data in a SQLite DB.
type Persistent struct {
	*Market
	db *sql.DB
}

// LoadFromFile opens path as a SQLite database and loads market data.
func LoadFromFile(path string) *Persistent {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Printf("db open error: %v", err)
		return &Persistent{Market: New()}
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS tokens(address TEXT PRIMARY KEY);`); err != nil {
		log.Printf("token table: %v", err)
	}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS pools(address TEXT PRIMARY KEY, token0 TEXT, token1 TEXT);`); err != nil {
		log.Printf("pool table: %v", err)
	}
	p := &Persistent{Market: New(), db: db}

	rows, err := db.Query(`SELECT address FROM tokens`)
	if err == nil {
		for rows.Next() {
			var a string
			if rows.Scan(&a) == nil {
				p.Market.AddToken(common.HexToAddress(a))
			}
		}
		rows.Close()
	}
	rows, err = db.Query(`SELECT address, token0, token1 FROM pools`)
	if err == nil {
		for rows.Next() {
			var addr, t0, t1 string
			if rows.Scan(&addr, &t0, &t1) == nil {
				p.Market.AddPool(common.HexToAddress(addr), common.HexToAddress(t0), common.HexToAddress(t1))
			}
		}
		rows.Close()
	}
	return p
}

func (p *Persistent) saveToken(addr common.Address) {
	if p.db != nil {
		_, _ = p.db.Exec(`INSERT OR IGNORE INTO tokens(address) VALUES (?)`, addr.Hex())
	}
}

func (p *Persistent) savePool(addr, t0, t1 common.Address) {
	if p.db != nil {
		_, _ = p.db.Exec(`INSERT OR IGNORE INTO pools(address, token0, token1) VALUES (?,?,?)`, addr.Hex(), t0.Hex(), t1.Hex())
	}
}

// Add records a pool address with no token metadata.
func (p *Persistent) Add(addr common.Address) {
	if !p.Market.Has(addr) {
		p.Market.Add(addr)
		if p.db != nil {
			_, _ = p.db.Exec(`INSERT OR IGNORE INTO pools(address) VALUES (?)`, addr.Hex())
		}
	}
}

// AddPool records a pool and saves to the DB.
func (p *Persistent) AddPool(addr, token0, token1 common.Address) {
	if !p.Market.Has(addr) {
		p.Market.AddPool(addr, token0, token1)
		p.savePool(addr, token0, token1)
	}
}

// AddToken records a token and saves to the DB.
func (p *Persistent) AddToken(addr common.Address) {
	if !p.Market.HasToken(addr) {
		p.Market.AddToken(addr)
		p.saveToken(addr)
	}
}

func (p *Persistent) Close() error {
	if p.db != nil {
		return p.db.Close()
	}
	return nil
}

// PoolsForTokens returns pools that trade the given token pair.
func (p *Persistent) PoolsForTokens(t0, t1 common.Address) []common.Address {
	return p.Market.PoolsForTokens(t0, t1)
}
