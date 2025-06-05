package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    "math"
)

type OpenSkyResponse struct {
    Time uint64                 `json:"time"`
    States [][]interface{}      `json:"states"`
}

func fetchOpenSkyData(o *OpenSkyResponse) {
    url := "https://opensky-network.org/api/states/all"

    response, err := http.Get(url)
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        fmt.Printf("Error: %v", err)
        return
    }

    err = json.Unmarshal(body, &o)

    if err != nil {
        fmt.Printf("Error: %v", err)
        return 
    }

}

func haversing(long1 float64, lat1 float64, long2 float64, lat2 float64) float64 {
    // radius of the earth
    R := 6371000

    phi_1 := lat1 * math.Pi/180.0
    phi_2 := lat1 * math.Pi/180.0

    delta_phi := (lat2 - lat1) * (math.Pi/180.0)
    delta_lambda := (long2 - long1) * (math.Pi/180.0)

    a := math.Pow(math.Sin(delta_phi / 2.0), 2) + math.Cos(phi_1) * math.Cos(phi_2) * math.Pow(math.Sin(delta_lambda / 2.0), 2)

    c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1 - a))

    meters := float64(R) * c
    return meters
}

func main() {
    long, lat := -117.2284, 32.8732
    test_long, test_lat := -117.2284, 32.8742


    var o OpenSkyResponse
    fetchOpenSkyData(&o)

    dist := haversing(long, lat, test_long, test_lat)
    fmt.Printf("Distance Test: %f\n", dist)

    for _, state := range(o.States) {
        origin_c := state[2]
        f_long := state[5]
        f_lat := state[6]
        fmt.Printf("Origin Country: %s\t%f, %f\n", origin_c, f_long, f_lat)
    }
}
