package api

import (
    "github.com/timkraemer1/flight-tracker/models"
)

func fetchStates() (models.OpenSkyResponse, error) {
    // TODO: fetch states api
    return models.OpenSkyResponse{}, nil
}

func fetchFlights() ([]models.FlightData, error) {
    // TODO: fetch flight api
    return nil, nil
}
