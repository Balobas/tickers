package fetcher

import (
	"encoding/json"
	"log"
	"net/http"
	"tickers/app/data"
	"tickers/app/interfaces/database"
	"tickers/app/repo"
	"time"
)

func Fetch(db database.Database, remoteUrl string, stop <-chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			resp, err := http.Get(remoteUrl)
			if err != nil {
				log.Println("Error: ", err)
				continue
			}

			var tickers []data.Ticker

			if err := json.NewDecoder(resp.Body).Decode(&tickers); err != nil {
				log.Println("Error: ", err)
				continue
			}

			if err := repo.NewTickerRepository(db).Store(tickers); err != nil {
				log.Println("Error: ", err)
				continue
			}

			time.Sleep(30 * time.Second)
		}
	}
}
