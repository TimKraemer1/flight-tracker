package main

import (
	"fmt"
	"log"
	"github.com/timkraemer1/flight-tracker/api"
	// "github.com/timkraemer1/flight-tracker/models"
	// "github.com/timkraemer1/flight-tracker/utils"
)

func main() {
	token, err := api.RetrieveAuthToken()
    if err != nil {
        fmt.Printf("%v\n", err)
        return
    } 

	err = api.FetchFlight(token, "a290d0")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}