package saver

import (
	"encoding/json"
	"fmt"
	"log"
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

var updatePriceQueryAVFunc string = "TIME_SERIES_DAILY"
var symbolSuffix string = ".BSE"

func fetchBSEData(symbol string) {
	symbolUppr := strings.ToUpper(symbol)
	finalUrl := AV_BaseUrl + updatePriceQueryAVFunc + "&symbol=" + symbolUppr + symbolSuffix + "&outputsize=compact" + "&apikey=" + AV_KEY
	log.Println(finalUrl)
	resp, err := http.Get(finalUrl)
	if err != nil {
		log.Printf(err.Error())

	}
	defer resp.Body.Close()
	log.Printf(resp.Status)
	var avResponse AvResponse
	json.NewDecoder(resp.Body).Decode(&avResponse)
	for k, v := range avResponse.TimeSeries {
		fmt.Println(k, v.Close)
	}

}
