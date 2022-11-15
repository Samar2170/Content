package saver

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/samar2170/Content/db"
)

func SaveDJI() error {
	resp, err := getTWDEOD("DJI")
	if err != nil {
		return err
	}
	index, err := db.GetIndexByCode("DJI")
	closeFloat, err := strconv.ParseFloat(resp.Close, 64)
	if err != nil {
		return err
	}
	indexData := db.IndexData{
		IndexId: index.ID,
		Date:    resp.Datetime,
		Close:   closeFloat,
	}
	return indexData.GetorCreate()
}

func SaveNYA() error {
	resp, err := getTWDEOD("NYA")
	if err != nil {
		return err
	}
	index, err := db.GetIndexByCode("NYA")
	closeFloat, err := strconv.ParseFloat(resp.Close, 64)
	if err != nil {
		return err
	}
	indexData := db.IndexData{
		IndexId: index.ID,
		Date:    resp.Datetime,
		Close:   closeFloat,
	}
	return indexData.GetorCreate()
}

func SaveIXIC() error {
	resp, err := getTWDEOD("ixic")
	if err != nil {
		return err
	}
	index, err := db.GetIndexByCode("IXIC")
	closeFloat, err := strconv.ParseFloat(resp.Close, 64)
	if err != nil {
		return err
	}
	indexData := db.IndexData{
		IndexId: index.ID,
		Date:    resp.Datetime,
		Close:   closeFloat,
	}
	return indexData.GetorCreate()
}

func SaveInvestingIndices() {
	resp, err := http.Get(INV_BaseUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	log.Print(resp.StatusCode)
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		panic(err)
	}
	item := doc.Find("div.table-browser_table-browser-wrapper__fhiVh")

	item.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
		symbol := s.Find("td").Eq(1).Text()
		price := s.Find("td").Eq(2).Text()
		// REMOVE COMMAS, HYPHENS, AND PERCENTAGES USING REGEX
		symbolF := regexp.MustCompile(`[^\w\s]`).ReplaceAllString(symbol, "")

		co, err := db.GetIndexByName(symbolF)

		if err != nil {
			log.Println(err)
			return
		}
		price = strings.Replace(price, ",", "", -1)
		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println(err)
			return
		}
		// time to string
		t, err := time.Now().UTC().MarshalText()
		if err != nil {
			log.Println(err)
			return
		}

		indexData := db.IndexData{
			IndexId: co.ID,
			Date:    string(t),
			Close:   priceFloat,
		}
		err = indexData.GetorCreate()
		if err != nil {
			log.Println(err)
		}
	})

}
