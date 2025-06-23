package utils

import (
	_ "github.com/mattn/go-sqlite3"
)

type SQLightFlightCache struct{
	db *sql.DB
	arrivalCacheSize int
	departureCacheSize int
}

func CreateSQLiteCache(dbPath string) (*SQLiteFlightCache, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS arrivalsCache (
			airport_code TEXT PRIMARY KEY,
			flights_json TEXT,
			cached_at INTEGER
		)
	`)

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS departuresCache (
			airport_code TEXT PRIMARY KEY,
			flights_json TEXT,
			cached_at INTEGER
		)
	`)
}