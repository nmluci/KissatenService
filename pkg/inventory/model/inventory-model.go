package models

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
)

const (
	INSERT_ITEM_STMT        = "INSERT INTO Inventory(name, price, stock) values (?, ?, ?)"
	COUNT_ITEM_STMT         = "SELECT COUNT(DISTINCT id) FROM Inventory"
	UPDATE_ITEM_STOCKS_STMT = "UPDATE Inventory SET stock = ? WHERE id == ?"
	UPDATE_ITEM_PRICE_STMT  = "UPDATE Inventory SET price = ? WHERE id == ?"
	UPDATE_ITEM_ALL_STMT    = "UPDATE Inventory SET price = ?, stock = ? WHERE id == ?"
	GET_ALL_ITEM_STMT       = "SELECT id, name, price, stock FROM Inventory"
	GET_ITEM_BY_ID_STMT     = "SELECT id, name, price, stock FROM Inventory WHERE id == ?"
	GET_ITEM_BY_QUERY       = "SELECT id, name, price, stock FROM Inventory WHERE name LIKE ?"
)

type InventoryModel struct {
	DB *sql.DB
}

type QueryData struct {
	Name       string `json:"name"`
	PriceStart string `json:"priceStart,omitempty"`
	PriceEnd   string `json:"priceEnd,omitempty"`
	StockStart string `json:"stockStart,omitempty"`
	StockEnd   string `json:"stockEnd,omitempty"`
}

type Item struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Items []*Item

func (im *InventoryModel) GetItemById(id int) (*Item, error) {
	rows, err := im.DB.Query(GET_ITEM_BY_ID_STMT, id)
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

func (im *InventoryModel) InsertItem(name string, price int, stocks int) (int, error) {
	rows, err := im.DB.Exec(INSERT_ITEM_STMT, name, price, stocks)
	if err != nil {
		log.Println(err)
		return -1, err
	}

	rowNum, err := rows.LastInsertId()
	if err != nil {
		log.Println(err)
		return -1, nil
	}
	return int(rowNum), nil
}

func (im *InventoryModel) GetItemByQuery(query *QueryData) (Items, error) {
	var items Items
	var itemVal []interface{}
	var baseQuery = GET_ITEM_BY_QUERY
	itemVal = append(itemVal, query.Name)

	if query.PriceStart != "" && query.PriceEnd != "" {
		baseQuery += "AND price BETWEEN ? AND ? "
		itemVal = append(itemVal, query.PriceStart, query.PriceEnd)
	}

	if query.StockStart != "" && query.StockEnd != "" {
		baseQuery += "AND stock BETWEEN ? AND ? "
		itemVal = append(itemVal, query.StockStart, query.StockEnd)
	}

	rows, err := im.DB.Query(baseQuery, itemVal...)
	if err != nil {
		log.Println(err)
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

func (item *Item) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(item)
}

func (item *Item) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(item)
}
