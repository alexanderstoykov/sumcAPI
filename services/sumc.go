package services

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type SumcProvider struct{}

type SumcResponse struct {
	Lines []struct {
		VehicleType string `json:"vehicle_type"`
		Arrivals    []struct {
			Time string `json:"time"`
		} `json:"arrivals"`
		Name string `json:"name"`
	} `json:"lines"`
}

func (*SumcProvider) CallAPI(busStop int) (bytes []byte) {
	client := http.Client{}
	url := "https://api-arrivals.sofiatraffic.bg/api/v1/arrivals/" + strconv.Itoa(busStop) + "/"

	resp, _ := client.Get(url)
	bytes, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	return
}
