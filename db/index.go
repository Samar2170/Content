package db

import "gorm.io/gorm"

type Index struct {
	*gorm.Model
	Name    string
	Code    string
	Country string
}

type IndexData struct {
	*gorm.Model
	IndexId uint
	Index   Index  `gorm:"foreignKey:IndexId"`
	Date    string `gorm:"Index"`
	Close   float64
}

func (i Index) GetorCreate() error {
	return db.FirstOrCreate(&i, Index{Name: i.Name, Code: i.Code, Country: i.Country}).Error
}
func (i IndexData) GetorCreate() error {
	return db.FirstOrCreate(&i, IndexData{IndexId: i.IndexId, Date: i.Date, Close: i.Close}).Error
}

func GetIndexByCode(code string) (Index, error) {
	var index Index
	err := db.Where("code = ?", code).First(&index).Error
	return index, err
}
func GetIndexByName(name string) (Index, error) {
	var index Index
	err := db.Where("name ILIKE ?", name).First(&index).Error
	return index, err
}

func FillIndices() {
	indices := []Index{
		{Name: "Dow Jones Industrial Average", Code: "DJI", Country: "USA"},
		{Name: "S&P 500", Code: "SPX", Country: "USA"},
		{Name: "NASDAQ Composite", Code: "IXIC", Country: "USA"},
		{Name: "NYSE COMPOSITE (DJ)", Code: "NYA", Country: "USA"},

		// Fetched via https://www.investing.com/indices/major-indices
		{Name: "Nifty 50", Code: "NIFTY 50", Country: "India"},
		{Name: "CSE All Share", Code: "CSE ALL", Country: "Sri Lanka"},
		{Name: "Bovespa", Code: "IBOV", Country: "Brazil"},
		{Name: "CAC 40", Code: "FCHI", Country: "France"},
		{Name: "IBEX 35", Code: "IBEX", Country: "Spain"},
		{Name: "FTSE 100", Code: "UKX", Country: "UK"},
		{Name: "Hang Seng", Code: "HSI", Country: "Hong Kong"},
		{Name: "KOSPI", Code: "KOSPI", Country: "South Korea"},
		{Name: "Nikkei 225", Code: "NKY", Country: "Japan"},
	}
	for _, index := range indices {
		index.GetorCreate()
	}
}
