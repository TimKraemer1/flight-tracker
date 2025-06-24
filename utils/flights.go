package utils

import (
	"fmt"
	"strings"
	"time"
	"sort"
	"github.com/timkraemer1/flight-tracker/models"
)

func FormatArrivals(arrivals []models.FlightData) string {
	if len(arrivals) == 0 {
		return "No arrival data available"
	}

	sort.Slice(arrivals, func(i, j int) bool {
		return arrivals[i].LastSeen < arrivals[j].LastSeen
	})

	var sb strings.Builder
	for i, flight := range arrivals {
		departureTime := time.Unix(flight.FirstSeen, 0).Format("03:04PM")
		arrivalTime := time.Unix(flight.LastSeen, 0).Format("03:04PM")

		sb.WriteString(fmt.Sprintf(
			"[green]Flight %d\n[white]Callsign: [yellow]%s\n[white]Departure: [cyan]%s [white] at [blue]%s\n[white]Arrival: [cyan]%s [white]at [blue]%s\n\n",
			i+1,
			strings.TrimSpace(flight.CallSign),
			flight.EstDepartureAirport,
			departureTime,
			flight.EstArrivalAirport,
			arrivalTime,
		))
	}

	return sb.String()
}

func FormatDepartures(departures []models.FlightData) string {
	if len(departures) == 0 {
		return "No departure data available"
	}

	sort.Slice(departures, func(i, j int) bool {
		return departures[i].FirstSeen < departures[j].FirstSeen
	})

	var sb strings.Builder
	for i, flight := range departures {
		departureTime := time.Unix(flight.FirstSeen, 0).Format("03:04PM")
		arrivalTime := time.Unix(flight.LastSeen, 0).Format("03:04PM")

		sb.WriteString(fmt.Sprintf(
			"[green]Flight %d\n[white]Callsign: [yellow]%s\n[white]Departure: [cyan]%s [white] at [blue]%s\n[white]Arrival: [cyan]%s [white]at [blue]%s\n\n",
			i+1,
			strings.TrimSpace(flight.CallSign),
			flight.EstDepartureAirport,
			departureTime,
			flight.EstArrivalAirport,
			arrivalTime,
		))
	}

	return sb.String()
}