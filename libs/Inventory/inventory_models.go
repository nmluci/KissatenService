package InventoryLibs

type Item struct {
	ID int
	name string
	sum int
	price uint
}

type Inventory []Item