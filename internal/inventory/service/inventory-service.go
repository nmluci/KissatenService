package service

import (
	models "github.com/nmluci/KissatenService/internal/inventory/model"
)

type Inventory map[string]*models.Item

var invFile = "./storage/inventory.data"

// Initialize the Inventory by retrieving data from the prefer database (PlainText for now),
// and load it into the Inventory Struct
