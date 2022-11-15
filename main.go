package main

import (
	"log"
	"os"
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
	log.Println(time.Now())

}

func main() {
	LOGFILE := "/tmp/Content.log"
	logFile, err := os.OpenFile(LOGFILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Day().At(DJ_ClosingTime).Do(func() {
		saveUSIndices()
	})
	s.Every(1).Day().At(NIFTY_ClosingTime).Do(func() {
		saver.SaveInvestingIndices()
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
	NIFTY_ClosingTime  = "10:01:00"
)
