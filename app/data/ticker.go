package data

type Ticker struct {
	Symbol         string  `json:"symbol"`
	Price24        float64 `json:"price_24h" gorm:"column:price_24h"`
	Volume24       float64 `json:"volume_24h" gorm:"column:volume_24h"`
	LastTradePrice float64 `json:"last_trade_price"`
}

type TickerRepository interface {
	GetAll() ([]Ticker, error)
	Store(tickers []Ticker) error
}
