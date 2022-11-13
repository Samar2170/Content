package saver

import (
	"strconv"

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
