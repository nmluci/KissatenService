package InventoryLibs

import (
	"fmt"
)

var InventoryList = Inventory{
	{0, "Cafe au Lait", 90, 4500},
	{1, "Cappucino", 50, 4500},
	{2, "Vanilla Latte", 70, 4500},
	{3, "Espresso", 80, 4500},
}

func ItemCount() int {
	return len(InventoryList)
}

func ItemByID(id int) *Item {
	for _, item := range InventoryList {
		if item.ID == id {
			return &item
		}
	}
	return nil
}

func PrintItem() {
	for _, item := range InventoryList {
		fmt.Printf("%s %d %d\n", item.name, item.sum, item.price)
	}
}
