package sqlite

import (
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	schema = `
CREATE TABLE IF NOT EXISTS tickers (
	symbol VARCHAR(32) PRIMARY KEY,
	price_24h FLOAT,
	volume_24h FLOAT,
	last_trade_price FLOAT

);

CREATE INDEX IF NOT EXISTS ticker_symbol ON tickers(symbol);
`
)

type SqliteDB struct {
	db *gorm.DB
}

func (s *SqliteDB) Gorm() *gorm.DB {
	return s.db
}

func NewDB(file string) (*SqliteDB, error) {

	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if err := db.Exec(schema); err.Error != nil {
		return nil, errors.WithStack(err.Error)
	}

	return &SqliteDB{db: db}, nil
}
