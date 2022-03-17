package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/saintmalik/ifihadinvested/api_services"
	"github.com/saintmalik/ifihadinvested/model"
)

var tpl *template.Template


func init() {
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}
func Ifihad(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		trending(w, r)
	case "/hodling":
		tpl.ExecuteTemplate(w, "hodling.html", nil)
	default:
		fmt.Fprintf(w, "You lost the way comrade!!")
	}
}

func trending(w http.ResponseWriter, r *http.Request) {

	trending_data_res_body, err := api_services.TrendingData(model.Id) //get trending data of coin
	if err != nil {
		log.Fatal("No response from request")
	}

	var readtrend model.CoinTrendingData
	err = json.Unmarshal(trending_data_res_body, &readtrend) // unmarshal JSON into Go value
	if err != nil {
		fmt.Println("error:", err)
	}

trend_coin0 := readtrend.Coins[0].Item.Name
trendcoin_rank0 := readtrend.Coins[0].Item.MarketCapRank
trendcoin_thumb0 := readtrend.Coins[0].Item.Thumb
trendcoin_id0 := readtrend.Coins[0].Item.ID
trendcoin_symbol0 := readtrend.Coins[0].Item.Symbol


trend_coin1 := readtrend.Coins[1].Item.Name
trendcoin_rank1 := readtrend.Coins[1].Item.MarketCapRank
trendcoin_thumb1 := readtrend.Coins[1].Item.Thumb
trendcoin_id1 := readtrend.Coins[1].Item.ID
trendcoin_symbol1 := readtrend.Coins[1].Item.Symbol

trend_coin2 := readtrend.Coins[2].Item.Name
trendcoin_rank2 := readtrend.Coins[2].Item.MarketCapRank
trendcoin_thumb2 := readtrend.Coins[2].Item.Thumb
trendcoin_id2 := readtrend.Coins[2].Item.ID
trendcoin_symbol2 := readtrend.Coins[2].Item.Symbol

trend_coin3 := readtrend.Coins[3].Item.Name
trendcoin_rank3 := readtrend.Coins[3].Item.MarketCapRank
trendcoin_thumb3 := readtrend.Coins[3].Item.Thumb
trendcoin_id3 := readtrend.Coins[3].Item.ID
trendcoin_symbol3 := readtrend.Coins[3].Item.Symbol

trend_coin4 := readtrend.Coins[4].Item.Name
trendcoin_rank4 := readtrend.Coins[4].Item.MarketCapRank
trendcoin_thumb4 := readtrend.Coins[4].Item.Thumb
trendcoin_id4 := readtrend.Coins[4].Item.ID
trendcoin_symbol4 := readtrend.Coins[4].Item.Symbol

trend_coin5 := readtrend.Coins[5].Item.Name
trendcoin_rank5 := readtrend.Coins[5].Item.MarketCapRank
trendcoin_thumb5 := readtrend.Coins[5].Item.Thumb
trendcoin_id5 := readtrend.Coins[5].Item.ID
trendcoin_symbol5 := readtrend.Coins[5].Item.Symbol

d := struct {
	TrendCoinName0 string
	TrendCoinName1 string
	TrendCoinName2 string
	TrendCoinName3 string
	TrendCoinName4 string
	TrendCoinName5 string
	TrendCoinRank0 int
	TrendCoinRank1 int
	TrendCoinRank2 int
	TrendCoinRank3 int
	TrendCoinRank4 int
	TrendCoinRank5 int
	TrendCoinThumb0 string
	TrendCoinThumb1 string
	TrendCoinThumb2 string
	TrendCoinThumb3 string
	TrendCoinThumb4 string
	TrendCoinThumb5 string
	TrendCoinId0 string
	TrendCoinId1 string
	TrendCoinId2 string
	TrendCoinId3 string
	TrendCoinId4 string
	TrendCoinId5 string
	TrendCoinSymbol0 string
	TrendCoinSymbol1 string
	TrendCoinSymbol2 string
	TrendCoinSymbol3 string
	TrendCoinSymbol4 string
	TrendCoinSymbol5 string

}{
	TrendCoinName0: trend_coin0,
	TrendCoinName1: trend_coin1,
	TrendCoinName2: trend_coin2,
	TrendCoinName3: trend_coin3,
	TrendCoinName4: trend_coin4,
	TrendCoinName5: trend_coin5,
	TrendCoinRank0: trendcoin_rank0,
	TrendCoinRank1: trendcoin_rank1,	
	TrendCoinRank2: trendcoin_rank2,
	TrendCoinRank3: trendcoin_rank3,
	TrendCoinRank4: trendcoin_rank4,
	TrendCoinRank5: trendcoin_rank5,
	TrendCoinThumb0: trendcoin_thumb0,
	TrendCoinThumb1: trendcoin_thumb1,
	TrendCoinThumb2: trendcoin_thumb2,
	TrendCoinThumb3: trendcoin_thumb3,
	TrendCoinThumb4: trendcoin_thumb4,
	TrendCoinThumb5: trendcoin_thumb5,
	TrendCoinId0: trendcoin_id0,
	TrendCoinId1: trendcoin_id1,
	TrendCoinId2: trendcoin_id2,
	TrendCoinId3: trendcoin_id3,
	TrendCoinId4: trendcoin_id4,
	TrendCoinId5: trendcoin_id5,
	TrendCoinSymbol0: trendcoin_symbol0,
	TrendCoinSymbol1: trendcoin_symbol1,
	TrendCoinSymbol2: trendcoin_symbol2,
	TrendCoinSymbol3: trendcoin_symbol3,
	TrendCoinSymbol4: trendcoin_symbol4,
	TrendCoinSymbol5: trendcoin_symbol5,
}

tpl.ExecuteTemplate(w, "index.html", d)

}
