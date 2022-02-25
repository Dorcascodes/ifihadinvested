package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"log"
	"strings"
	"io/ioutil"
	"strconv"
	"encoding/json"
)

var tpl *template.Template

func init(){
    tpl = template.Must(template.ParseGlob("templates/*.html",))
}

func ifihad(w http.ResponseWriter, r *http.Request){
    tpl.ExecuteTemplate(w, "form.html", nil)
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

var id string
var fromdate string
var fund string
var code string
var current_price string


func invested (w http.ResponseWriter, r *http.Request){

	fromdate = r.FormValue("from")
	id = r.FormValue("coinid")
	fund = r.FormValue("fiat_amount")

   res, err := http.Get("https://api.coingecko.com/api/v3/coins/"+id+"/history?date=" +fromdate+ "&localization=false")

		if err != nil {
			log.Fatal("No response from request")
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
buying_price := read.MarketData.CurrentPrice.Usd
latest_price := reading.MarketData.CurrentPrice.Usd

capital, err := strconv.ParseFloat(fund, 64)

total_owned := float64(capital) / buying_price

total_value_now := total_owned * latest_price

change_percent := (latest_price - buying_price) / buying_price * 100

profit_loss := total_value_now - capital

coin_symbol := strings.ToUpper(read.Symbol)

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
 InvestmentDate: fromdate,
}
// a struct to hold the data and pass them to the template to be deisplayed

tpl.ExecuteTemplate(w, "worthnow.html", k)
}
 // a switch case to check the country code to decide the number format stripping, so it runs a range over the numbers and strips the required things off based on the rules


// a struct to hold the data and pass them to the template to be deisplayed

// tpl.ExecuteTemplate(w, "index.html", nil)