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

// inisialisasi database dan auto migrasi ke mysql
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Bunker{})
}

// method membuat data bunker
func (b *Bunker) CreateBunker() *Bunker {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// method mengambil semua data bunker yang ada di database
func GetAllBunker() []Bunker {
	var Bunkers []Bunker
	db.Find(&Bunkers)
	return Bunkers
}

// method mengambil data bunker berdasarkan id
func GetBunkerById(Id int64) (*Bunker, *gorm.DB) {
	var getBunker Bunker
	db := db.Where("ID=?", Id).Find(&getBunker)
	return &getBunker, db
}

// method menghapus data bunker
func DeleteBunker(ID int64) Bunker {
	var bunker Bunker
	db.Where("ID=?", ID).Delete(bunker)
	return bunker
}
