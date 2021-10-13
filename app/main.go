package main

import (
	"github.com/pkg/errors"
	"log"
	"tickers/app/interfaces/database"
	"tickers/app/interfaces/fetcher"
	"tickers/app/interfaces/httpServer"
)

const remoteUrl = "https://api.blockchain.com/v3/exchange/tickers"

func main() {
	db, err := database.NewSQLiteDB("/tmp/database")
	if err != nil {
		log.Fatal(errors.Wrap(err, "cant init database"))
	}

	config := httpServer.Config{
		Port: "3000",
		Db:   db,
	}

	stopFetch := make(chan bool, 2)

	go fetcher.Fetch(db, remoteUrl, stopFetch)
	defer func() {
		stopFetch <- true
		close(stopFetch)
	}()

	httpServer.Run(config)
}
