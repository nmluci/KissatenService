package service

import (
	"errors"
	"log"

	models "github.com/nmluci/KissatenService/internal/inventory/model"
)

func GetAllItem(im *models.InventoryModel) (models.Items, error) {
	if data, err := im.GetAllItem(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func GetItemByName(im *models.InventoryModel, name string) (*models.Item, error) {
	if data, err := im.GetItemByName(name); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func GetItemById(im *models.InventoryModel, id int) (*models.Item, error) {
	if data, err := im.GetItemById(id); err != nil {
		return nil, err
	} else {
		return data, err
	}
}

func RemoveItem(im *models.InventoryModel, itemId int) error {
	// Considered a Best Practiced to just removed its stock instead of remove it from the database
	if err := im.UpdateItemAllProp(itemId, -1, -1); err != nil {
		return err
	} else {
		return nil
	}
}

func InsertItem(im *models.InventoryModel, name string, price int, stock int) (*int, error) {
	if id, err := im.InsertItem(name, price, stock); err != nil {
		log.Println(err)
		return nil, err
	} else {
		return &id, nil
	}
}

func UpdateItem(im *models.InventoryModel, itemId int, newItem *models.Item) error {
	if itemId != newItem.Id {
		return errors.New("changing item id is not permitted. considered adding them as a new item instead")
	}

	if err := im.UpdateItemAllProp(newItem.Id, newItem.Stock, newItem.Price); err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateItemPrice(im *models.InventoryModel, itemId int, newPrice int) error {
	if newPrice < -1 {
		return errors.New("price must be greater than 0")
	}
	if err := im.UpdateItemPrice(itemId, newPrice); err != nil {
		return err
	} else {
		return nil
	}
}

func UpdateItemStocks(im *models.InventoryModel, itemid int, newStocks int) error {
	if err := im.UpdateItemStocks(itemid, newStocks); err != nil {
		return err
	} else {
		return nil
	}
}
