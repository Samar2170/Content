package saver

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type TWDResponse struct {
	MetaData struct {
		Symbol           string `json:"symbol"`
		Interval         string `json:"interval"`
		Currency         string `json:"currency"`
		ExchangeTimezone string `json:"exchange_timezone"`
		Exchange         string `json:"exchange"`
		MicCode          string `json:"mic_code"`
		Type             string `json:"type"`
	} `json:"meta"`
	Values []struct {
		Datetime string `json:"datetime"`
		Open     string `json:"open"`
		High     string `json:"high"`
		Low      string `json:"low"`
		Close    string `json:"close"`
		Volume   int64  `json:"volume"`
	} `json:"values"`
}

func getTWDDataTS(symbol string) (TWDResponse, error) {
	symbol = strings.ToUpper(symbol)
	finalUrl := TWD_BaseUrl + "time_series?symbol=" + symbol + "&interval=1day&apikey=" + TWD_KEY
	log.Println(finalUrl)
	resp, err := http.Get(finalUrl)
	if err != nil {
		return TWDResponse{}, err
	}
	defer resp.Body.Close()
	var twdResponse TWDResponse
	json.NewDecoder(resp.Body).Decode(&twdResponse)
	return twdResponse, nil

}

type TWDResponseEOD struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	MicCode  string `json:"mic_code"`
	Currency string `json:"currency"`
	Datetime string `json:"datetime"`
	Close    string `json:"close"`
}

func getTWDEOD(symbol string) (TWDResponseEOD, error) {
	symbol = strings.ToUpper(symbol)
	finalUrl := TWD_BaseUrl + "eod?symbol=" + symbol + "&apikey=" + TWD_KEY
	log.Println(finalUrl)
	resp, err := http.Get(finalUrl)
	if err != nil {
		return TWDResponseEOD{}, err
	}
	defer resp.Body.Close()
	var twdResponse TWDResponseEOD
	json.NewDecoder(resp.Body).Decode(&twdResponse)
	return twdResponse, nil
}

type TWDResponseExchangeRate struct {
	Symbol    string  `json:"symbol"`
	Rate      float64 `json:"rate"`
	Timestamp uint    `json:"timestamp"`
}

func getTWDExchangeRate(symbol string) (TWDResponseExchangeRate, error) {
	symbol = strings.ToUpper(symbol)
	finalUrl := TWD_BaseUrl + "exchange_rate?symbol=" + symbol + "&apikey=" + TWD_KEY
	log.Println(finalUrl)
	resp, err := http.Get(finalUrl)
	if err != nil {
		return TWDResponseExchangeRate{}, err
	}
	defer resp.Body.Close()
	var twdResponse TWDResponseExchangeRate

	json.NewDecoder(resp.Body).Decode(&twdResponse)
	return twdResponse, nil
}
