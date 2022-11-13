package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/samar2170/Content/saver"
)

func saveUSIndices() {
	saver.SaveDJI()
	saver.SaveNYA()
	saver.SaveIXIC()
}

func printTime() {
	fmt.Println(time.Now())

}

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At(DJ_ClosingTime).Do(func() {
		saveUSIndices()
	})
	s.Every(1).Day().At(FOREX_ARBIT_TIME).Do(func() {
		saver.SaveForexPairs()
	})

	s.Every("1m").Do(func() {
		printTime()
	})
	s.StartAsync()
	s.StartBlocking()

}

const (
	DJ_ClosingTime     = "21:01:00"
	NYSE_ClosingTime   = "21:01:00"
	NASDAQ_ClosingTime = "21:01:00"
	FOREX_ARBIT_TIME   = "22:01:00"
)
