package models

type OpenSkyResponse struct {
    Time        uint64          `json:"time"`
    States      [][]interface{} `json:"states"`
}
