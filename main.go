package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"log"
	"strings"
	"io/ioutil"
    "time"
	"strconv"
	"encoding/json"
)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("templates/*.html",))
}

func ifihad(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "index.html", nil)
}

func holdling(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "hodling.html", "nil")
}

func main() {

	port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
    }

    http.HandleFunc("/", ifihad)
	http.HandleFunc("/worthnow", invested)
    http.HandleFunc("/hodling", holdling)
    http.HandleFunc("/hodl", ifihadhodl)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    fmt.Printf("Starting server for testing HTTP POST...\n")
    if err := http.ListenAndServe(":"+port, nil); err != nil {
        log.Fatal(err)
    }
}

var (
 id string
 fromdate string
 fund string
 historical_data string
 current_price string
 check_response string
 crypto_coin string
 expected_amount string
)

func ifihadhodl (w http.ResponseWriter, r *http.Request){

	fromdate = r.FormValue("from")
	id = r.FormValue("coinid")
	crypto_coin = r.FormValue("coinowned") //picking up the value from the form
    expected_amount = r.FormValue("fiat_expeectation")

   res, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"/history?date=" +fromdate+ "&localization=false") //get historical data of coin

		if err != nil {
			fmt.Println("No response from request")
		}

   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body) // response body is []byte
   historical_data = fmt.Sprint(string(body)) 

   type CoinHistoricalData struct {
    ID string `json:"id"`
    Symbol string `json:"symbol"`
    Name   string `json:"name"`
    MarketData struct {
        CurrentPrice struct {
            Usd  float64 `json:"usd"`
            Sats float64 `json:"sats"`
        } `json:"current_price"`
    } `json:"market_data"`
}   //struct to store data

jsonString := historical_data 
var read CoinHistoricalData
err = json.Unmarshal([]byte(jsonString), &read) //Unmarshal JSON to struct

if err != nil {
    fmt.Println("error:", err)
}

coin_owned, err := strconv.ParseFloat(crypto_coin, 64)

amoount_you_feel_worth_in_some_day, err := strconv.ParseFloat(expected_amount, 64)

coin_value_on_sell_date := read.MarketData.CurrentPrice.Usd

ifyousold := coin_value_on_sell_date * coin_owned

ifyouhodl := coin_owned * amoount_you_feel_worth_in_some_day

profit_loss := ifyouhodl - ifyousold

coin_symbol := strings.ToUpper(read.Symbol)

// change_percentage := (ifyousold/ifyouhodl)*100

dateparsing := fromdate
layout := "01-01-2006"
t, _ := time.Parse(layout, dateparsing)
dateparsed := t.Format("01 Jan, 2006") 

IfHodlDatas := struct {
    IfHodl float64
    IfYouSold float64
    ProfitLoss float64
    CoinSymbol string
    CoinValueOnSellDate float64
    CoinOwned float64
    AmountYouFeelWorthInSomeDay float64
    // ChangePercentage float64
    SellingDate string
}{
    IfHodl: ifyouhodl,
    IfYouSold: ifyousold,
    ProfitLoss: profit_loss,
    CoinSymbol: coin_symbol,
    CoinValueOnSellDate: coin_value_on_sell_date,
    CoinOwned: coin_owned,
    AmountYouFeelWorthInSomeDay: amoount_you_feel_worth_in_some_day,
    // ChangePercentage: change_percentage,
    SellingDate: dateparsed,
}

fmt.Println(IfHodlDatas.CoinValueOnSellDate)
    tpl.ExecuteTemplate(w, "hodl.html", IfHodlDatas)
}


func invested (w http.ResponseWriter, r *http.Request){

	fromdate = r.FormValue("from")
	id = r.FormValue("coinid")
	fund = r.FormValue("fiat_amount") //picking up the value from the form

   res, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"/history?date=" +fromdate+ "&localization=false") //get historical data of coin

		if err != nil {
			fmt.Println("No response from request")
		}

   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body) // response body is []byte
   historical_data = fmt.Sprint(string(body)) 

   type CoinHistoricalData struct {
    ID string `json:"id"`
    Symbol string `json:"symbol"`
    Name   string `json:"name"`
    MarketData struct {
        CurrentPrice struct {
            Usd  float64 `json:"usd"`
            Sats float64 `json:"sats"`
        } `json:"current_price"`
    } `json:"market_data"`
}   //struct to store data

jsonString := historical_data 
var read CoinHistoricalData
err = json.Unmarshal([]byte(jsonString), &read) //Unmarshal JSON to struct

if err != nil {
    fmt.Println("error:", err)
}

rep, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false") // make a GET request to take current price

if err != nil {
    log.Fatal("No response from request")
   }
   defer rep.Body.Close()
   body1, err := ioutil.ReadAll(rep.Body) // response body is []byte
   current_price = fmt.Sprint(string(body1))
//    fmt.Println(PrettyString(code))

type CoinCurrentData struct {
    ID              string `json:"id"`
    Symbol          string `json:"symbol"`
    MarketData          struct {
        CurrentPrice struct {
            Usd  float64 `json:"usd"`
            Sats float64 `json:"sats"`
        } `json:"current_price"`
    } `json:"market_data"`
} //struct for the json data

jsonCurrent := current_price

var reading CoinCurrentData
err = json.Unmarshal([]byte(jsonCurrent), &reading) // unmarshal JSON into Go value

if err != nil {
    fmt.Println("error:", err)
}

if read.MarketData.CurrentPrice.Usd != 0 {
    fmt.Sprint("Nice")
} else {
    check_response = fmt.Sprint("No data found")
}

buying_price := read.MarketData.CurrentPrice.Usd

latest_price := reading.MarketData.CurrentPrice.Usd

capital, err := strconv.ParseFloat(fund, 64) //convert string value to float

total_owned := float64(capital) / buying_price

total_value_now := total_owned * latest_price

change_percent := (latest_price - buying_price) / buying_price * 100

profit_loss := total_value_now - capital

coin_symbol := strings.ToUpper(read.Symbol)

dateparsing := fromdate
layout := "01-01-2006"
t, _ := time.Parse(layout, dateparsing)
dateparsed := t.Format("01 Jan, 2006") //date parsing and formatting

CoinDatas := struct {
	CoinName string
	CoinSymbol string
	LatestCoinPrice float64
	TotalCoinOwned float64
	TotalValueNow float64
	PurchasePrice float64
	ChangePercent float64
	ProfitLoss float64
	InitialCapital float64
	InvestmentDate string
    CheckResponse string
}{
 CoinName: read.Name,
 CoinSymbol: coin_symbol,
 LatestCoinPrice: latest_price,
 TotalCoinOwned: total_owned,
 TotalValueNow: total_value_now,
 PurchasePrice: buying_price,
 ChangePercent: change_percent,
 ProfitLoss: profit_loss,
 InitialCapital: capital,
 InvestmentDate: dateparsed,
CheckResponse: check_response,
}

tpl.ExecuteTemplate(w, "worthnow.html", CoinDatas)
}