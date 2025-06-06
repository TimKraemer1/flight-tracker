package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    "math"
    "github.com/timkraemer1/flight-tracker/api"
    "github.com/timkraemer1/flight-tracker/models"
)


func main() {
    token := api.retrieveAuthToken()
    fmt.Printf("Token: %s\n", token)
}
