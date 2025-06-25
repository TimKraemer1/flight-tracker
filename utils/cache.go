package utils

import (
	"database/sql"
	"time"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"github.com/timkraemer1/flight-tracker/models"
	"github.com/timkraemer1/flight-tracker/api"
)

type SQLiteFlightCache struct{
    db *sql.DB
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
	if err != nil {
		return nil, err
	}

    _, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS departuresCache (
            airport_code TEXT PRIMARY KEY,
            flights_json TEXT,
            cached_at INTEGER
        )
    `)
	if err != nil {
		return nil, err
	}

	return &SQLiteFlightCache{db: db}, nil
}

func (s *SQLiteFlightCache) LoadArrivalsFromCache(airportCode string, maxAge time.Duration) ([]models.FlightData, bool, error){
	row := s.db.QueryRow(`SELECT flights_json, cached_at FROM arrivalsCache WHERE airport_code = ?`, airportCode)
	var jsonData string
	var cachedInt int64

	err := row.Scan(&jsonData, cachedInt)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, false, nil
		}
		return nil, false, err
	}

	cachedAt := time.Unix(cachedInt, 0)
	if time.Since(cachedAt) > maxAge {
		return nil, false, nil
	}

	var flights []models.FlightData
	err = json.Unmarshal([]byte(jsonData), &flights)
	if err != nil {
		return nil, false, err
	}

	return flights, true, nil
}

func (s *SQLiteFlightCache) SaveArrivalsToCache(airportCode string, flights []models.FlightData) error {
	jsonData, err := json.Marshal(flights)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(`INSERT OR REPLACE INTO arrivalsCache (airport_code, flights_json, cached_at) VALUES (?, ?, ?)`, airportCode, string(jsonData), time.Now().Unix())
	return err
}

func (s *SQLiteFlightCache) GetArrivals(token, airport string) ([]models.FlightData, error) {
	maxAge := 24 * time.Hour

	flights, hit, err := s.LoadArrivalsFromCache(airport, maxAge)
	if err != nil {
		return nil, err
	}

	if hit {
		return flights, nil
	}

	flights, err = api.FetchArrivals(token, airport)
	if err != nil {
		return nil, err
	}

	err = s.SaveArrivalsToCache(airport, flights)
	if err != nil {
		return nil, err
	}

	return flights, nil
}