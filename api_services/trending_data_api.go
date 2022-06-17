package api_services
import (
	"fmt"
	"io/ioutil"
	"net/http"
)
func TrendingData(id string) ([]byte, error) {
	//get historical data of coin
	res, err := http.Get("https://api.coingecko.com/api/v3/search/trending")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer res.Body.Close()
	trending_data_body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error:", err)
	}
	return trending_data_body, err
}
