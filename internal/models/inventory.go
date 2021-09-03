package models

type Item struct {
	Id    int
	Name  string
	Sum   int
	Price uint
}

type Inventory map[string]*Item
