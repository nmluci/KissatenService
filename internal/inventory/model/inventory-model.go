package models

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
)

const (
	INSERT_ITEM_STMT        = "INSERT INTO Inventory(id, name, price, stock) values (?, ?, ?, ?)"
	UPDATE_ITEM_STOCKS_STMT = "UPDATE Inventory SET stock = ? WHERE id == ?"
	UPDATE_ITEM_PRICE_STMT  = "UPDATE Inventory SET price = ? WHERE id == ?"
	UPDATE_ITEM_ALL_STMT    = "UPDATE Inventory SET price = ?, stock = ? WHERE id == ?"
	GET_ALL_ITEM_STMT       = "SELECT id, name, price, stock FROM Inventory"
	GET_ITEM_BY_ID_STMT     = "SELECT id, name, price, stock FROM Inventory WHERE id == ?"
	GET_ITEM_ID_STMT        = "SELECT id FROM Inventory WHERE name == ?"
)

type InventoryModel struct {
	DB *sql.DB
}

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Items []*Item

func (im *InventoryModel) GetItemByName(name string) (*Item, error) {
	rows, err := im.DB.Query(GET_ITEM_BY_ID_STMT, name)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	itm := &Item{}
	defer rows.Close()
	rows.Next()
	rows.Scan(&itm.Id, &itm.Name, &itm.Price, &itm.Stock)
	return itm, nil
}

func (im *InventoryModel) GetAllItem() (Items, error) {
	var items Items
	rows, err := im.DB.Query(GET_ALL_ITEM_STMT)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp Item
		rows.Scan(&temp.Id, &temp.Name, &temp.Price, &temp.Stock)
		items = append(items, &temp)
	}
	return items, nil
}

func (im *InventoryModel) AppendItem(newItem Item) error {
	if _, err := im.DB.Exec(INSERT_ITEM_STMT, newItem.Id, newItem.Name, newItem.Price, newItem.Stock); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (im *InventoryModel) UpdateItemPrice(itemId int, price int) error {
	if _, err := im.DB.Exec(UPDATE_ITEM_PRICE_STMT, price, itemId); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (im *InventoryModel) UpdateItemStocks(itemId int, stocks int) error {
	if _, err := im.DB.Exec(UPDATE_ITEM_STOCKS_STMT, stocks, itemId); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (im *InventoryModel) UpdateItemAllProp(itemId int, price int, stocks int) error {
	if _, err := im.DB.Exec(UPDATE_ITEM_ALL_STMT, price, stocks, itemId); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (im *InventoryModel) GetItemID(name string) (int, error) {
	if rows, err := im.DB.Query(GET_ITEM_ID_STMT, name); err != nil {
		return -1, err
	} else {
		var itemId int
		defer rows.Close()
		rows.Next()
		rows.Scan(&itemId)
		return itemId, nil
	}
}

func (item *Item) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(item)
}

func (item *Item) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(item)
}
