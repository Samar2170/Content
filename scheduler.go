package main

import (
	"fmt"
	"time"

	"github.com/samar2170/Content/saver"
)

func saveUSIndices() {
	saver.SaveDJI()
	saver.SaveNYA()
	saver.SaveIXIC()
}

func testCron() {
	fmt.Println("testCron")
	fmt.Println("testCron")
}

func printTime() {
	fmt.Println(time.Now())

}

// func main() {
// 	s := gocron.NewScheduler(time.UTC)

// 	s.Every(1).Day().At(DJ_ClosingTime).Do(func() {
// 		saveUSIndices()
// 	})

// 	s.Every("1m").Do(func() {
// 		printTime()
// 	})
// 	s.StartAsync()
// 	s.StartBlocking()

// }

func main() {
	saver.SaveForexPairs()
	// db.FillForexPairs()
}

// in UTC
const (
	DJ_ClosingTime     = "21:01:00"
	NYSE_ClosingTime   = "21:01:00"
	NASDAQ_ClosingTime = "21:01:00"
)
