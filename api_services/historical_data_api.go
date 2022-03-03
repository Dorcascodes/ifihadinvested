package api_services

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HistoricalData(id, fromdate string) ([]byte, error) {
	//get historical data of coin
	res, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id + "/history?date=" + fromdate + "&localization=false")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	historical_data_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	return historical_data_body, err
}
