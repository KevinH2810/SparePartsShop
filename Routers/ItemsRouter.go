package Routers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	. "SparePartsShop/db"
	"SparePartsShop/model"
	. "SparePartsShop/objects"

	"github.com/gorilla/mux"
)

type dbConn struct {
	DB *sql.DB
}

var (
	items []Item
)

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db, err := sql.Open("sqlite3", "sparepartshop.db")
	defer db.Close()
	feed := model.NewFeed(db)

	item := feed.GetItems()
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	db, err := sql.Open("sqlite3", "sparepartshop.db")
	defer db.Close()
	feed := model.NewFeed(db)

	item := feed.GetItem(id)
	json.NewEncoder(w).Encode(item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(item)
	// item.ID = strconv.Itoa(rand.Intn(1000000))
	//ID SHOULD BE HANDLED BY DB LATER
	items = append(items, item)
	json.NewEncoder(w).Encode(&item)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if id, _ := strconv.Atoi(params["id"]); item.ID == id {
			items = append(items[:index], items[index+1:]...)
			var post Item
			_ = json.NewDecoder(r.Body).Decode(post)
			item.ID, _ = strconv.Atoi(params["id"])
			items = append(items, post)
			json.NewEncoder(w).Encode(&post)
			return
		}
	}
	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if id, err := strconv.Atoi(params["id"]); item.ID == id {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
