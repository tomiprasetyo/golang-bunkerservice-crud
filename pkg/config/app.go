package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// global variable
var (
	db *gorm.DB
)

// buat koneksi ke database
func Connect() {
	d, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/go_bunkerservice_crud?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}

	db = d
}

// ambil database
func GetDB() *gorm.DB {
	return db
}
