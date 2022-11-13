package saver

import (
	"fmt"
	"time"

	"github.com/samar2170/Content/db"
)

func SaveForexPairs() {
	indices, err := db.GetForexPairs()
	if err != nil {
		return
	}
	for _, index := range indices {
		// fmt.Println(index)
		// fmt.Println(index.Id)

		twdresp, err := getTWDExchangeRate(index.Symbol)
		if err != nil {
			fmt.Println(err)
			continue
		}
		datetime := time.Unix(int64(twdresp.Timestamp), 0)
		indexData := db.ForexData{
			ForexPairId: index.Id,
			Date:        datetime,
			Close:       twdresp.Rate,
		}
		err = indexData.GetorCreate()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(25 * time.Second)
	}
}
