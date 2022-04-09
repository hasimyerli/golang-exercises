package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	getBinanceTickers()

	getBtcTurkTickers()
}

func getBinanceTickers() {
	fmt.Println("Start Binance")

	type BinanceTicker []struct {
		Symbol             string `json:"symbol"`
		PriceChange        string `json:"priceChange"`
		PriceChangePercent string `json:"priceChangePercent"`
	}

	res, _ := http.Get("https://api2.binance.com/api/v3/ticker/24hr")

	tickerData, _ := ioutil.ReadAll(res.Body)

	var tickers BinanceTicker

	json.Unmarshal(tickerData, &tickers)

	data, _ := json.MarshalIndent(tickers, "", " ")

	ioutil.WriteFile("binanceTiclers.json", data, 0644)

	fmt.Println("Finish Binance")
}

func getBtcTurkTickers() {
	fmt.Println("Start BtcTurk")

	type BtcTurkTicker struct {
		Data []struct {
			Pair         string  `json:"pair"`
			Daily        float32 `json:"daily"`
			DailyPercent float32 `json:"dailyPercent"`
		} `json:"data"`
	}

	res, _ := http.Get("https://api.btcturk.com/api/v2/ticker")

	tickerData, _ := ioutil.ReadAll(res.Body)

	var tickers BtcTurkTicker

	json.Unmarshal(tickerData, &tickers)

	data, _ := json.MarshalIndent(tickers.Data, "", " ")

	ioutil.WriteFile("btcTurkTiclers.json", data, 0644)

	fmt.Println("Finish BtcTurk")
}
