package models

import (
	"database/sql"
)

type KissatenModel struct {
	DB *sql.DB
}

type UserCart struct {
	OrderId  int `json:"order_id"`
	MemberId int `json:"member_id"`
	ItemId   int `json:"item_id"`
	Sum      int `json:"sum"`
}
