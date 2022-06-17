package api_services
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func CurrentData(id string) ([]byte, error) {
	//get historical data of coin
	res, err := http.Get("https://api.coingecko.com/api/v3/coins/" + id + "?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	current_data_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	return current_data_body, err
}
