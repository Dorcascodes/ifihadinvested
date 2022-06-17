package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/saintmalik/ifihadinvested/api_services"
	"github.com/saintmalik/ifihadinvested/model"
)

func Compare(w http.ResponseWriter, r *http.Request) {
	model.Id = r.FormValue("coinid")
	model.Vscoinid = r.FormValue("vscoinid")
	current_data_res_body, err := api_services.CurrentData(model.Id) //get current data of coin
	if err != nil {
		fmt.Println("No response from request", err)
	}
	var reading model.CoinCurrentData
	err = json.Unmarshal(current_data_res_body, &reading) // unmarshal JSON into Go value
	if err != nil {
		fmt.Println("error reading:", err)
	}
	vscurrent_data_res_body, err := api_services.CurrentData(model.Vscoinid) //get current data of coin
	if err != nil {
		fmt.Println("No response from request", err)
	}
	var reader model.CoinCurrentData
	err = json.Unmarshal(vscurrent_data_res_body, &reader) // unmarshal JSON into Go value
	if err != nil {
		fmt.Println("error reader:", err)
	}

	coin_symbol := strings.ToUpper(reading.Symbol)
	vscoin_symbol := strings.ToUpper(reader.Symbol)

	latest_price := reading.MarketData.CurrentPrice.Usd
	// new_price := read.MarketCapRank
	compareDatas := model.CompareDatas{
		CoinName:          reading.Name,
		CoinSymbol:        coin_symbol,
		LatestCoinPrice:   latest_price,
		CoinRank:          reading.MarketCapRank,
		AllTimeHigh:       reading.MarketData.Ath.Usd,
		AllTimeHighDate:   reading.MarketData.AthDate.Usd,
		VsCoinName:        reader.Name,
		VsCoinSymbol:      vscoin_symbol,
		VsLatestCoinPrice: reader.MarketData.CurrentPrice.Usd,
		VsCoinRank:        reader.MarketCapRank,
		VsAllTimeHigh:     reader.MarketData.Ath.Usd,
		VsAllTimeHighDate: reader.MarketData.AthDate.Usd,
	}


	err = tpl.ExecuteTemplate(w, "compare.html", compareDatas)
	if err != nil {
		fmt.Println("Template error:", err)
	}
}
