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

	flights, err := api.FetchArrivals(token, "KSFO")
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	fmt.Print(flights)
}