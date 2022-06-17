package handlers
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/saintmalik/ifihadinvested/api_services"
	"github.com/saintmalik/ifihadinvested/model"
)
func Invested(w http.ResponseWriter, r *http.Request) {
	model.Fromdate = r.FormValue("from")
	model.Id = r.FormValue("coinid")
	model.Fund = r.FormValue("fiat_amount") //picking up the value from the form
	historical_res_body, err := api_services.HistoricalData(model.Id, model.Fromdate) //get historical data of coin
	if err != nil {
		log.Println("error:", err)
	}
	var read model.CoinHistoricalData
	err = json.Unmarshal(historical_res_body, &read) //Unmarshal JSON to struct
	if err != nil {
		fmt.Println("error:", err)
	}
	dateparsing := model.Fromdate
	t, _ := time.Parse(model.LayoutISO, dateparsing)
	dateparsed := t.Format(model.LayoutUS) //date parsing and formatting
	var emptyMarketData model.MarketData
	var coinDatas model.CoinDatas
	if read.MarketData == emptyMarketData {
		model.Check_response = "No data found"
		coinDatas.CheckResponse = model.Check_response
		coinDatas.InvestmentDate = dateparsed
		coinDatas.CoinSymbol = strings.ToUpper(read.Symbol)
		coinDatas.CoinName = read.Name
		tpl.ExecuteTemplate(w, "worthnow.html", coinDatas)
		return //returns if no data found
	}
	// fmt.Println("*******Historical Data:******", read.MarketData.CurrentPrice.Usd) //debugging
	current_data_res_body, err := api_services.CurrentData(model.Id) //get current data of coin
	if err != nil {
		log.Fatal("No response from request")
	}
	var reading model.CoinCurrentData
	err = json.Unmarshal(current_data_res_body, &reading) // unmarshal JSON into Go value
	if err != nil {
		fmt.Println("error:", err)
	}
	buying_price := read.MarketData.CurrentPrice.Usd
	latest_price := reading.MarketData.CurrentPrice.Usd
	capital, err := strconv.ParseFloat(model.Fund, 64) //convert string value to float
	if err != nil {
		fmt.Println("string conver error", err)
	}
	total_owned := float64(capital) / buying_price
	total_value_now := total_owned * latest_price
	change_percent := (latest_price - buying_price) / buying_price * 100
	profit_loss := total_value_now - capital
	coin_symbol := strings.ToUpper(read.Symbol)
	switch {
	case profit_loss > 0:
		model.Check_profit_loss = "This is a profit of"
	case profit_loss < 0:
		model.Check_profit_loss = "This is a loss of"
	} //profit and loss check
	coinDatas = model.CoinDatas{
		CoinName:             read.Name,
		CoinSymbol:           coin_symbol,
		LatestCoinPrice:      latest_price,
		TotalCoinOwned:       total_owned,
		TotalValueNow:        total_value_now,
		PurchasePrice:        buying_price,
		ChangePercent:        change_percent,
		ProfitLoss:           profit_loss,
		InitialCapital:       capital,
		InvestmentDate:       dateparsed,
		CheckResponse:        "",
		CheckValueProfitLoss: model.Check_profit_loss,
	}
	err = tpl.ExecuteTemplate(w, "worthnow.html", coinDatas)
	if err != nil {
		fmt.Println("Template error:", err)
	}
}
