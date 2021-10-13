package converter

import (
	"tickers/app/data"
)

func ConvertTickers(tickers []data.Ticker) []map[string]map[string]interface{} {
	var result []map[string]map[string]interface{}

	for _, ticker := range tickers {
		out := make(map[string]map[string]interface{})

		out[ticker.Symbol] = make(map[string]interface{})

		out[ticker.Symbol]["price_24h"] = ticker.Price24
		out[ticker.Symbol]["volume_24h"] = ticker.Volume24
		out[ticker.Symbol]["last_trade_price"] = ticker.LastTradePrice

		result = append(result, out)
	}

	return result
}
