package Routers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"SparePartsShop/model"
	. "SparePartsShop/objects"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

type (
	dbConn struct {
		DB *sql.DB
	}
)

var (
	items   []Item
	ItemObj Item
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//INJECT TECHNIQUE
	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	feed := model.NewFeed(db)

	//INITIATE DATABASE
	item, err := feed.GetItems()
	if err != nil {
		fmt.Println("GetItems = ", err)
	}
	json.NewEncoder(w).Encode(&item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	feed := model.NewFeed(db)

	item, err := feed.GetItem(params["id"])
	if err != nil {
		fmt.Println("GetItem = ", err)
	}
	json.NewEncoder(w).Encode(&item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&ItemObj)

	if err != nil {
		fmt.Println("Error = ", err)
		return
	}

	fmt.Println("itemObj Raw =", ItemObj.Company)

	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	feed := model.NewFeed(db)

	item, err := feed.CreateItem(ItemObj.Nama, ItemObj.Code, ItemObj.Jenis, ItemObj.Company, ItemObj.BuyPrice, ItemObj.SellPrice1, ItemObj.SellPrice2)
	if err != nil {
		fmt.Println("CreateItem = ", err)
	}
	json.NewEncoder(w).Encode(&item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var itemSlice []Item
	err := json.NewDecoder(r.Body).Decode(&itemSlice)

	if err != nil {
		fmt.Println("Error = ", err)
		return
	}

	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	feed := model.NewFeed(db)

	item, err := feed.UpdateItem(itemSlice)
	if err != nil {
		fmt.Println("UpdateItem = ", err)
	}
	json.NewEncoder(w).Encode(&item)
}

//WORKS BUT WE STILL HAVE TO THROW THE ENTIRE OBJECT INCLUDING OTHER USELESS DATA
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var itemSlice []Item
	_ = json.NewDecoder(r.Body).Decode(&itemSlice)

	db, err := sql.Open("sqlite3", "./db/sparepartshop.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	feed := model.NewFeed(db)
	fmt.Println(itemSlice)

	item, err := feed.DeleteItem(itemSlice)
	if err != nil {
		fmt.Println("DeleteItem = ", err)
	}
	json.NewEncoder(w).Encode(&item)
}
