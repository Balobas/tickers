package repo

import (
	"github.com/pkg/errors"
	"tickers/app/data"
	"tickers/app/interfaces/database"
)

type Ticker struct {
	Db database.Database
}

func NewTickerRepository(db database.Database) data.TickerRepository {
	return &Ticker{Db: db}
}

func (t Ticker) GetAll() ([]data.Ticker, error) {
	var tickers []data.Ticker
	res := t.Db.Gorm().Find(&tickers)
	if res.Error != nil {
		return nil, errors.WithStack(res.Error)
	}

	return tickers, nil
}

func (t Ticker) Store(tickers []data.Ticker) error {
	res := t.Db.Gorm().Save(tickers)
	if res.Error != nil {
		return errors.WithStack(res.Error)
	}
	return nil
}
