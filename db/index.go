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
	Index   Index `gorm:"foreignKey:IndexId"`
	Date    string
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

func FillIndices() {
	indices := []Index{
		{Name: "Dow Jones Industrial Average", Code: "DJI", Country: "USA"},
		{Name: "S&P 500", Code: "SPX", Country: "USA"},
		{Name: "NASDAQ Composite", Code: "IXIC", Country: "USA"},
		{Name: "NYSE COMPOSITE (DJ)", Code: "NYA", Country: "USA"},
	}
	for _, index := range indices {
		index.GetorCreate()
	}
}
