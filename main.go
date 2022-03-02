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

func main() {

	port := os.Getenv("PORT")
    if port == "" {
      port = "8080"
    }

    http.HandleFunc("/", ifihad)
	http.HandleFunc("/worthnow", invested)
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
 code string
 current_price string
 check_response string
)

func invested (w http.ResponseWriter, r *http.Request){

	fromdate = r.FormValue("from")
	id = r.FormValue("coinid")
	fund = r.FormValue("fiat_amount")

   res, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"/history?date=" +fromdate+ "&localization=false")

		if err != nil {
			fmt.Println("No response from request")
		}

   defer res.Body.Close()
   body, err := ioutil.ReadAll(res.Body) // response body is []byte
   code = fmt.Sprint(string(body)) 

   type d struct {
    ID string `json:"id"`
    Symbol string `json:"symbol"`
    Name   string `json:"name"`
    MarketData struct {
        CurrentPrice struct {
            Usd  float64 `json:"usd"`
            Sats float64 `json:"sats"`
        } `json:"current_price"`
    } `json:"market_data"`
}

jsonString := code
var read d
err = json.Unmarshal([]byte(jsonString), &read)

if err != nil {
    fmt.Println("error:", err)
}

rep, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"?localization=false&tickers=false&market_data=true&community_data=false&developer_data=false&sparkline=false")

if err != nil {
    log.Fatal("No response from request")
   }
   defer rep.Body.Close()
   body1, err := ioutil.ReadAll(rep.Body) // response body is []byte
   current_price = fmt.Sprint(string(body1))
//    fmt.Println(PrettyString(code))

type p struct {
    ID              string `json:"id"`
    Symbol          string `json:"symbol"`
    MarketData          struct {
        CurrentPrice struct {
            Usd  float64 `json:"usd"`
            Sats float64 `json:"sats"`
        } `json:"current_price"`
    } `json:"market_data"`
}

jsonCurrent := current_price

var reading p
err = json.Unmarshal([]byte(jsonCurrent), &reading)

if err != nil {
    fmt.Println("error:", err)
}

pricing := read.MarketData.CurrentPrice.Usd
fmt.Println(pricing)

// if pricing != 0 {
//     fmt.Sprint("Nice")
// } else {
//     look_response = "No data found"
// }

// if read.MarketData.CurrentPrice.Usd != 0 {
//     fmt.Println("Current price: ", read.MarketData.CurrentPrice.Usd)
// } else {
//     check_response = fmt.Sprint("No data found")
// }

buying_price := read.MarketData.CurrentPrice.Usd
latest_price := reading.MarketData.CurrentPrice.Usd

capital, err := strconv.ParseFloat(fund, 64)

total_owned := float64(capital) / buying_price

total_value_now := total_owned * latest_price

change_percent := (latest_price - buying_price) / buying_price * 100

profit_loss := total_value_now - capital

coin_symbol := strings.ToUpper(read.Symbol)

const (
    layoutISO = "02-01-2006"
    layoutUS  = "January 2, 2006"
)
dateparsing := fromdate
t, _ := time.Parse(layoutISO, dateparsing)
dateparsed := t.Format(layoutUS)

k := struct {
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
    CheckResponse float64
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
CheckResponse: pricing,
}

tpl.ExecuteTemplate(w, "worthnow.html", k)
}