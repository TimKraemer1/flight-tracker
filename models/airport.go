package models

type Airport struct {
	Icao 		string 	`json:"icao"`
	Name 		string 	`json:"name"`
	City 		string 	`json:"city"`
	State 		string 	`json:"state"`
	Country 	string 	`json:"country"`
	Elevation 	int 	`json:"elevation"`
	Latitude 	float32 `json:"lat"`
	Longitude 	float32 `json:"lon"`
}