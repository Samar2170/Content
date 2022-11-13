package saver

import (
	"encoding/json"
	"net/http"
	"strings"
)

type AvResponse struct {
	MetaData struct {
		Information string `json:"1. Information"`
		Symbol      string `json:"2. Symbol"`
		LastRefresh string `json:"3. Last Refreshed"`
		OutputSize  string `json:"5. Output Size"`
		TimeZone    string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	TimeSeries map[string]struct {
		Open   string `json:"1. open"`
		High   string `json:"2. high"`
		Low    string `json:"3. low"`
		Close  string `json:"4. close"`
		Volume string `json:"5. volume"`
	} `json:"Time Series (Daily)"`
}

var updatePriceQueryAVFunc string = "TIME_SERIES_DAILY_ADJUSTED"
var symbolSuffix string = ".BSE"

func FetchAVData(symbol string) (AvResponse, error) {
	symbolUppr := strings.ToUpper(symbol)
	finalUrl := AV_BaseUrl + updatePriceQueryAVFunc + "&symbol=" + symbolUppr + "&outputsize=compact" + "&apikey=" + AV_KEY
	resp, err := http.Get(finalUrl)
	if err != nil {
		return AvResponse{}, err
	}
	defer resp.Body.Close()

	var avResponse AvResponse
	json.NewDecoder(resp.Body).Decode(&avResponse)
	return avResponse, nil
}

func FetchBSEData(symbol string) (AvResponse, error) {
	symbolUppr := strings.ToUpper(symbol)
	resp, err := FetchAVData(symbolUppr + symbolSuffix)
	if err != nil {
		return AvResponse{}, err
	}
	return resp, nil
}
