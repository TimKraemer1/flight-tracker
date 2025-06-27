package utils

import (
	"encoding/json"
	"os"
	"fmt"
	"sync"
	"github.com/timkraemer1/flight-tracker/models"
)

var airportFileMutex sync.Mutex

func LoadAirportData(filename string) (map[string]models.Airport, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var airportMap map[string]models.Airport
	if err := json.Unmarshal(data, &airportMap); err != nil {
		return nil, err
	}

	return airportMap, nil
}

func AirportExists(filename, code string) (bool, models.Airport, error) {
	airportFileMutex.Lock()
	defer airportFileMutex.Unlock()

	file, err := os.Open(filename)
	if err != nil {
		return false, models.Airport{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	t, err := decoder.Token()
	if err != nil {
		return false, models.Airport{}, err
	}

	delim, ok := t.(json.Delim)
	if !ok || delim != '{' {
		return false, models.Airport{}, fmt.Errorf("expected start of json object")
	}

	var airport models.Airport
	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return false, models.Airport{}, err
		}

		key := t.(string)
		if key == code {
			err := decoder.Decode(&airport)
			if err != nil {
				return false, models.Airport{}, err
			}
			return true, airport, nil
		} else {
			var skip interface{}
			_ = decoder.Decode(&skip)
		}
	}
	return false, models.Airport{}, nil
}

func FormatAirportInfo(airport models.Airport) string {
	// Create a nicely formatted airport information display
	info := fmt.Sprintf(`[white::b]╔══════════════════════════════════════════════════════════════╗
║                     [yellow::b]AIRPORT INFORMATION[white::b]                      ║
╚══════════════════════════════════════════════════════════════╝

  [cyan::b]Airport Name:[white::-]  %-42s  
  [cyan::b]ICAO Code:[white::-]     %-42s  
                                                              
╔══════════════════════════════════════════════════════════════╗
║                        [green::b]LOCATION DETAILS[white::b]                      ║
╚══════════════════════════════════════════════════════════════╝
                                                             
  [magenta::b]City:[white::-]         %-42s  
  [magenta::b]State/Region:[white::-] %-42s  
  [magenta::b]Country:[white::-]      %-42s  
                                                              
╔══════════════════════════════════════════════════════════════╗
║                      [blue::b]TECHNICAL INFORMATION[white::b]                   ║
╚══════════════════════════════════════════════════════════════╝
                                                              
  [red::b]Elevation:[white::-]    %d ft                    
  [red::b]Longitude:[white::-]    %f°                              
  [red::b]Latitude:[white::-]     %f°                              
                                                              

[yellow::b]Navigation:[white::-] Press [green::b]'b' to go back to previous page`,
		airport.Name,
		airport.Icao,
		airport.City,
		airport.State,
		airport.Country,
		airport.Elevation,
		airport.Longitude,
		airport.Latitude)

	return info
}