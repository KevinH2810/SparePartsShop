package model

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	. "SparePartsShop/objects"

	"github.com/gorilla/mux"
)

var (
	items      []Item
	id         int
	Name       string
	Code       string
	Jenis      string
	Company    string
	BuyPrice   int
	SellPrice1 int
	SellPrice2 int
)

type (
	Feed struct {
		DB *sql.DB
	}
)

func NewFeed(db *sql.DB) *Feed {
	return &Feed{
		DB: db,
	}
}

//GetItems : Get All Items on database regardless of its type
func (feed *Feed) GetItems() []Item {
	rows, _ := feed.DB.Query("Select id, Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2 from Item")

	for rows.Next() {
		rows.Scan(&id, &Name, &Code, &Jenis, &Company, &BuyPrice, &SellPrice1, &SellPrice2)
		items = append(items, Item{ID: id, Nama: Name, Code: Code, Jenis: Jenis, CompanyManufacturer: Company, BuyPrice: BuyPrice, SellPrice1: SellPrice1, SellPrice2: SellPrice2})

	}
	return items
}

//GetItem : Get a single item based on supplied id
func (feed *Feed) GetItem(ids string) []Item {
	rows, _ := feed.DB.Query("Select id, Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2 from Item")

	for rows.Next() {
		rows.Scan(&id, &Name, &Code, &Jenis, &Company, &BuyPrice, &SellPrice1, &SellPrice2)
		if idInput, _ := strconv.Atoi(ids); id == idInput {
			items = append(items, Item{ID: id, Nama: Name, Code: Code, Jenis: Jenis, CompanyManufacturer: Company, BuyPrice: BuyPrice, SellPrice1: SellPrice1, SellPrice2: SellPrice2})
		}

	}
	// for _, item := range items {
	// 	if id, _ := strconv.Atoi(params["id"]); item.ID == id {
	// 		json.NewEncoder(w).Encode(item)
	// 		break
	// 	}
	// 	return
	// }
	return items
}

func (feed *Feed) CreateItem() {
	statement, _ := feed.DB.Prepare("Insert Into Item (Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2) Values(?,?,?,?,?,?,?)")
	statement.Exec()

	var item Item
	_ = json.NewDecoder(r.Body).Decode(item)
	// item.ID = strconv.Itoa(rand.Intn(1000000))
	//ID SHOULD BE HANDLED BY DB LATER
	items = append(items, item)
	json.NewEncoder(w).Encode(&item)
}

func (feed *Feed) UpdateItem(w http.ResponseWriter, r *http.Request) {
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

func (feed *Feed) DeleteItem(w http.ResponseWriter, r *http.Request) {
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
