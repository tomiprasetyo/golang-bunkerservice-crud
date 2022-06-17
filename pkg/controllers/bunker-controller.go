package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/models"
	"github.com/tomiprasetyo/golang-bunkerservice-crud/pkg/utils"
)

var NewBunker models.Bunker

// mengambil data bunker melalui models
func GetBunker(w http.ResponseWriter, r *http.Request) {
	newBunkers := models.GetAllBunker()
	res, _ := json.Marshal(newBunkers)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// mengambil data bunker melalui models berdasarkan id
func GetBunkerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bunkerId := vars["bunkerId"]
	ID, err := strconv.ParseInt(bunkerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bunkerDetails, _ := models.GetBunkerById(ID)
	res, _ := json.Marshal(bunkerDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// membuat data bunker dan menyimpannya ke dalam database
func CreateBunker(w http.ResponseWriter, r *http.Request) {
	CreateBunker := &models.Bunker{}
	utils.ParseBody(r, CreateBunker)
	b := CreateBunker.CreateBunker()
	res, _ := json.Marshal(b)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// menghapus data bunker berdasarkan id
func DeleteBunker(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bunkerId := vars["bunkerId"]
	ID, err := strconv.ParseInt(bunkerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bunker := models.DeleteBunker(ID)
	res, _ := json.Marshal(bunker)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBunker(w http.ResponseWriter, r *http.Request) {
	var updateBunker = &models.Bunker{}
	utils.ParseBody(r, updateBunker)

	vars := mux.Vars(r)
	bunkerId := vars["bunkerId"]
	ID, err := strconv.ParseInt(bunkerId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bunkerDetails, db := models.GetBunkerById(ID)
	if bunkerDetails.NoSO != "" {
		bunkerDetails.NoSO = updateBunker.NoSO
	}

	if bunkerDetails.NamaPerusahaan != "" {
		bunkerDetails.NamaPerusahaan = updateBunker.NamaPerusahaan
	}

	if bunkerDetails.NamaKapal != "" {
		bunkerDetails.NamaKapal = updateBunker.NamaKapal
	}

	if bunkerDetails.Product != "" {
		bunkerDetails.Product = updateBunker.Product
	}

	if bunkerDetails.Quantity != "" {
		bunkerDetails.Quantity = updateBunker.Quantity
	}

	if bunkerDetails.Pelabuhan != "" {
		bunkerDetails.Pelabuhan = updateBunker.Pelabuhan
	}

	db.Save(&bunkerDetails)
	res, _ := json.Marshal(bunkerDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
