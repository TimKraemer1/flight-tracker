package models

type FlightData struct {
    Icao24                  string  `json:"icao24"`
    EstDepartureAirport     string  `json:"estDepartureAirport"`
    EstArrivalAirport       string  `json:"estArrivalAirport"`
    FirstSeen               int64   `json:"firstSeen"`
    LastSeen                int64   `json:"lastSeen"`
    CallSign                string  `json:"callsign"`
}
