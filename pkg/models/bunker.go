package models

import (
	"github.com/jinzhu/gorm"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/config"
)

var db *gorm.DB

type Bunker struct {
	gorm.Model
	NoSO           string `json:"noSO"`
	NamaPerusahaan string `json:"namaPerusahaan"`
	NamaKapal      string `json:"namaKapal"`
	Product        string `json:"product"`
	Quantity       string `json:"quantity"`
	Pelabuhan      string `json:"pelabuhan"`
}

// inisialisasi database dan auto migrasi ke mysql
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Bunker{})
}

// method membuat data bunker di database
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

// method menghapus data bunker yang ada di database
func DeleteBunker(ID int64) Bunker {
	var bunker Bunker
	db.Where("ID=?", ID).Delete(bunker)
	return bunker
}
