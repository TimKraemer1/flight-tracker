package utils

import (
	"fmt"
	"strings"
	"time"
	"sort"
	"github.com/timkraemer1/flight-tracker/models"
)

// Format arrival data from models.FlightData -> string
func FormatArrivals(arrivals []models.FlightData) string {
	if len(arrivals) == 0 {
		return "No arrival data available"
	}

	// Sort arrivals by time
	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].LastSeen < arrivals[j].LastSeen
	})

	airportMap, err := LoadAirportData("airports.json")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var sb strings.Builder
	sb.WriteString("[yellow::b]Navigation:[white::-] Press [green::b]'b' [white::-]to go back to previous page\n\n")

	for i, flight := range arrivals {
		callSign := strings.TrimSpace(flight.CallSign)
		departureTime := time.Unix(flight.FirstSeen, 0).Format("01/02 03:04PM MST")
		arrivalTime := time.Unix(flight.LastSeen, 0).Format("01/02 03:04PM MST")

		estDep := flight.EstDepartureAirport
		locationDep := ""
		if airport, ok := airportMap[estDep]; ok {
			estDep = airport.Name
			locationDep = fmt.Sprintf(" - %s %s, %s", airport.City, airport.State, airport.Country)
		} else {
			estDep = "[red]Unknown"
		}

		estArr := flight.EstArrivalAirport
		locationArr := ""
		if airport, ok := airportMap[estArr]; ok {
			estArr = airport.Name
			locationArr = fmt.Sprintf(" - %s %s, %s", airport.City, airport.State, airport.Country)
		} else {
			estArr = "[red]Unknown"
		}

		sb.WriteString(fmt.Sprintf(
			"[green]Flight %d\n[white]Callsign: [yellow]%s, %s\n[white]Departure: [cyan]%s%s\n[blue]%s\n\n[white]Arrival: [cyan]%s%s\n[blue]%s\n\n",
			i+1,
			callSign,
			flight.Icao24,
			estDep,
			locationDep,
			departureTime,
			estArr,
			locationArr,
			arrivalTime,
		))
	}

	return sb.String()
}

// Format departure data from models.FlightData -> string
func FormatDepartures(departures []models.FlightData) string {
	if len(departures) == 0 {
		return "No departure data available"
	}

	// Sort departures by time
	sort.Slice(departures, func(i, j int) bool {
		return departures[i].LastSeen < departures[j].LastSeen
	})

	airportMap, err := LoadAirportData("airports.json")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	var sb strings.Builder
	sb.WriteString("[yellow::b]Navigation:[white::-] Press [green::b]'b' [white::-]to go back to previous page\n\n")

	for i, flight := range departures {
		callSign := strings.TrimSpace(flight.CallSign)
		departureTime := time.Unix(flight.FirstSeen, 0).Format("01/02 03:04PM MST")
		arrivalTime := time.Unix(flight.LastSeen, 0).Format("01/02 03:04PM MST")

		estDep := flight.EstDepartureAirport
		locationDep := ""
		if airport, ok := airportMap[estDep]; ok {
			estDep = airport.Name
			locationDep = fmt.Sprintf(" - %s %s, %s", airport.City, airport.State, airport.Country)
		} else {
			estDep = "[red]Unknown"
		}

		estArr := flight.EstArrivalAirport
		locationArr := ""
		if airport, ok := airportMap[estArr]; ok {
			estArr = airport.Name
			locationArr = fmt.Sprintf(" - %s %s, %s", airport.City, airport.State, airport.Country)
		} else {
			estArr = "[red]Unknown"
		}

		sb.WriteString(fmt.Sprintf(
			"[green]Flight %d\n[white]Callsign: [yellow]%s\n[white]Departure: [cyan]%s%s\n[blue]%s\n\n[white]Arrival: [cyan]%s%s\n[blue]%s\n\n",
			i+1,
			callSign,
			estDep,
			locationDep,
			departureTime,
			estArr,
			locationArr,
			arrivalTime,
		))
	}

	return sb.String()
}