package api

import (
    "encoding/json"
    "io/ioutil"
    "io"
    "net/http"
    "time"
    "fmt"
    "github.com/timkraemer1/flight-tracker/models"
)

func FetchStates(o *models.OpenSkyResponse) (error) {
    url := "https://opensky-network.org/api/states/all"

    response, err := http.Get(url)
    if err != nil {
        return err
    }
    defer response.Body.Close()

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return err
    }

    err = json.Unmarshal(body, o)
    if err != nil {
        return err
    }
    return nil
}

func FetchArrivals(token string, airport string) ([]models.FlightData, error) {
    now := time.Now().Unix()
    begin := now - (3600 * 24)
    end := now

    // Encode query parameters in the URL
    url := fmt.Sprintf(
        "https://opensky-network.org/api/flights/arrival?airport=%s&begin=%d&end=%d",
        airport, begin, end,
    )

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("unexpected status: %s, body: %s", resp.Status, string(body))
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var flights []models.FlightData
    err = json.Unmarshal(body, &flights)
    if err != nil {
        return nil, err
    }

    return flights, nil
}

func FetchDepartures(token string, airport string) ([]models.FlightData, error) {
        now := time.Now().Unix()
    begin := now - (3600 * 24)
    end := now

    // Encode query parameters in the URL
    url := fmt.Sprintf(
        "https://opensky-network.org/api/flights/departure?airport=%s&begin=%d&end=%d",
        airport, begin, end,
    )

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, err
    }

    req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return nil, fmt.Errorf("unexpected status: %s, body: %s", resp.Status, string(body))
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var flights []models.FlightData
    err = json.Unmarshal(body, &flights)
    if err != nil {
        return nil, err
    }

    return flights, nil
}