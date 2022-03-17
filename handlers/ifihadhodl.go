package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/saintmalik/ifihadinvested/api_services"
	"github.com/saintmalik/ifihadinvested/model"
)

func Ifihadhodl(w http.ResponseWriter, r *http.Request) {

	model.Fromdate = r.FormValue("from")
	model.Id = r.FormValue("coinid")
	model.Crypto_coin = r.FormValue("coinowned") //picking up the value from the form
	model.Expected_amount = r.FormValue("fiat_expectation")

	historical_data_res_body, err := api_services.HistoricalData(model.Id, model.Fromdate) //get current data of coin
	if err != nil {
		fmt.Println("No response from request")
	}

	var read model.CoinHistoricalData
	err = json.Unmarshal(historical_data_res_body, &read) //Unmarshal JSON to struct
	if err != nil {
		http.Error(w, "An unexpected error has occurred", http.StatusInternalServerError)
	}

	dateparsing := model.Fromdate
	t, _ := time.Parse(model.LayoutISO, dateparsing)
	dateparsed := t.Format(model.LayoutUS) //date parsing and formatting

	var emptyMarketData model.MarketData
	var IfHodlDatas model.IfHodlDatas
	if read.MarketData == emptyMarketData {
		model.Check_response = "No data found"
		IfHodlDatas.CheckResponse = model.Check_response
		IfHodlDatas.SellingDate = dateparsed
		IfHodlDatas.CoinSymbol = strings.ToUpper(read.Symbol)
		IfHodlDatas.CoinName = read.Name
		tpl.ExecuteTemplate(w, "hodl.html", IfHodlDatas)
		return //returns if no data found
	}

	coin_owned, err := strconv.ParseFloat(model.Crypto_coin, 64)
	if err != nil {
		fmt.Println("error:", err)
	}

	amount_you_feel_worth_in_some_day, err := strconv.ParseFloat(model.Expected_amount, 64)
	if err != nil {
		fmt.Println("error:", err)
	}

	coin_value_on_sell_date := read.MarketData.CurrentPrice.Usd

	ifyousold := coin_value_on_sell_date * coin_owned

	ifyouhodl := coin_owned * amount_you_feel_worth_in_some_day

	profit_loss := ifyouhodl - ifyousold

	coin_symbol := strings.ToUpper(read.Symbol)

	// change_percentage := (ifyousold/ifyouhodl)*100

	IfHodlDatas = model.IfHodlDatas {
		IfHodl:     ifyouhodl,
		IfYouSold:     ifyousold,
		ProfitLoss:       profit_loss,
		CoinSymbol:     coin_symbol,
		CoinValueOnSellDate:   coin_value_on_sell_date,
		CoinOwned:   coin_owned,
		AmountYouFeelWorthInSomeDay: amount_you_feel_worth_in_some_day,
		SellingDate: dateparsed,
	}

	tpl.ExecuteTemplate(w, "hodl.html", IfHodlDatas)
}
