package models

import (
	inv "github.com/nmluci/KissatenService/internal/Inventory"
	mem "github.com/nmluci/KissatenService/internal/Membership"
)

type UserShoppingCart struct {
	User *mem.Member
	Cart inv.Inventory
}

type UserList []*UserShoppingCart
