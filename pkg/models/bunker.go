package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/config"
)

var db *gorm.DB

type Bunker struct {
	gorm.Model
	ID             int    `gormjson:"id"`
	NoSO           string `json:"noSO"`
	NamaPerusahaan string `json:"namaPerusahaan"`
	NamaKapal      string `json:"namaKapal"`
	Quantity       string `json:"quantity"`
	Product        string `json:"product"`
	Pelabuhan      string `json:"pelabuhan"`
	Completed      bool   `json:"completed"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Bunker{})
}
