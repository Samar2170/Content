package db

import (
	"time"

	"gorm.io/gorm"
)

type ForexPair struct {
	*gorm.Model
	Id     uint `gorm:"primaryKey"`
	Base   string
	Quote  string
	Symbol string
}

type ForexData struct {
	*gorm.Model
	ForexPairId uint
	ForexPair   ForexPair `gorm:"foreignKey:ForexPairId"`
	Date        time.Time
	Close       float64
}

// func fillForexPairs()

func (f ForexPair) GetorCreate() error {
	return db.FirstOrCreate(&f, ForexPair{Base: f.Base, Quote: f.Quote, Symbol: f.Symbol}).Error
}

func (f ForexData) GetorCreate() error {
	return db.FirstOrCreate(&f, ForexData{ForexPairId: f.ForexPairId, Date: f.Date, Close: f.Close}).Error
}

func GetForexPairBySymbol(Symbol string) (ForexPair, error) {
	var forexPair ForexPair
	err := db.Where("Symbol = ?", Symbol).First(&forexPair).Error
	return forexPair, err
}

func GetForexPairs() ([]ForexPair, error) {
	var fps []ForexPair
	err := db.Find(&fps).Error
	return fps, err
}

func FillForexPairs() {
	pairs := []ForexPair{
		{Base: "USD", Quote: "CNH", Symbol: "USD/CNH"},
		{Base: "USD", Quote: "JPY", Symbol: "USD/JPY"},
		{Base: "USD", Quote: "EUR", Symbol: "USD/EUR"},
		{Base: "USD", Quote: "GBP", Symbol: "USD/GBP"},
		{Base: "USD", Quote: "AUD", Symbol: "USD/AUD"},
		{Base: "USD", Quote: "CAD", Symbol: "USD/CAD"},
		{Base: "USD", Quote: "CHF", Symbol: "USD/CHF"},
		{Base: "USD", Quote: "TRY", Symbol: "USD/TRY"},
		{Base: "USD", Quote: "SGD", Symbol: "USD/SGD"},
		{Base: "USD", Quote: "INR", Symbol: "USD/INR"},
		{Base: "USD", Quote: "IDR", Symbol: "USD/IDR"},
		{Base: "USD", Quote: "RUB", Symbol: "USD/RUB"},
		{Base: "USD", Quote: "BRL", Symbol: "USD/BRL"},
		{Base: "USD", Quote: "KRW", Symbol: "USD/KRW"},
	}
	for _, pair := range pairs {
		pair.GetorCreate()
	}

}
