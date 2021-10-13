package database

import "tickers/app/interfaces/database/sqlite"

func NewSQLiteDB(src string) (Database, error) {
	return sqlite.NewDB(src)
}
