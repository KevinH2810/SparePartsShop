package model

import (
	"database/sql"
	"fmt"
	"strconv"

	. "SparePartsShop/objects"
)

var (
	items                                []Item
	Name, Code, Jenis, Company           string
	id, BuyPrice, SellPrice1, SellPrice2 int
)

type (
	Feed struct {
		DB *sql.DB
	}
)

func NewFeed(db *sql.DB) *Feed {
	stmt, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS "Item" (
		"id"	INTEGER NOT NULL,
		"Name"	Varchar(100) NOT NULL,
		"Code"	Varchar(50),
		"Jenis"	Varchar(50),
		"Company"	Varchar(50),
		"BuyPrice"	INTEGER NOT NULL,
		"SellPrice1"	INTEGER,
		"SellPrice2"	INTEGER,
		PRIMARY KEY("id" AUTOINCREMENT)
	);`)

	stmt.Exec()
	return &Feed{
		DB: db,
	}
}

//GetItems : Get All Items on database regardless of its type
func (feed *Feed) GetItems() (*[]Item, error) {
	var item []Item
	rows, err := feed.DB.Query("Select id, Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2 from Item")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &Name, &Code, &Jenis, &Company, &BuyPrice, &SellPrice1, &SellPrice2)
		item = append(item, Item{ID: id, Nama: Name, Code: Code, Jenis: Jenis, Company: Company, BuyPrice: BuyPrice, SellPrice1: SellPrice1, SellPrice2: SellPrice2})
	}
	return &item, nil
}

//GetItem : Get a single item based on supplied id and return an array containing only 1 object
func (feed *Feed) GetItem(ids string) (*[]Item, error) {
	var item []Item
	rows, err := feed.DB.Query("Select id, Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2 from Item")

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		rows.Scan(&id, &Name, &Code, &Jenis, &Company, &BuyPrice, &SellPrice1, &SellPrice2)
		if idInput, _ := strconv.Atoi(ids); id == idInput {
			item = append(item, Item{ID: id, Nama: Name, Code: Code, Jenis: Jenis, Company: Company, BuyPrice: BuyPrice, SellPrice1: SellPrice1, SellPrice2: SellPrice2})
		}

	}

	return &item, nil
}

func (feed *Feed) CreateItem(Name string, Code string, Jenis string, Company string, BuyPrice int, SellPrice1 int, SellPrice2 int) (bool, error) {
	statement, _ := feed.DB.Prepare("Insert Into Item (Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2) Values(?,?,?,?,?,?,?)")
	_, err := statement.Exec(Name, Code, Jenis, Company, BuyPrice, SellPrice1, SellPrice2)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (feed *Feed) UpdateItem(item []Item) (bool, error) {

	statement, _ := feed.DB.Prepare(`UPDATE Item 
	SET name = ?,
	Code =?,
	Jenis =?,
	Company = ?,
	BuyPrice = ?,
	SellPrice1 = ?,
	SellPrice2 = ?
	Where id = ?`)

	for _, items := range item {
		_, err := statement.Exec(items.Nama, items.Code, items.Jenis, items.Company, items.BuyPrice, items.SellPrice1, items.SellPrice2, items.ID)

		if err != nil {
			fmt.Println("Error = ", err)
			return false, err
		}
	}
	return true, nil
}

func (feed *Feed) DeleteItem(Items []Item) (bool, error) {

	for _, item := range Items {
		statement, _ := feed.DB.Prepare(`DELETE FROM Item
		Where id = ?`)

		_, err := statement.Exec(item.ID)

		if err != nil {
			return false, err
		}
	}

	return true, nil
}
