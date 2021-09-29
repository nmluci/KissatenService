package models

import (
	"database/sql"
	"strconv"
	"time"
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

type orderMap map[int]int

var currentCarts orderMap

func (km *KissatenModel) MakeNewCart(userId int) int {
	orderId, ok := currentCarts[userId]
	if !ok {
		orderId, _ = strconv.Atoi(time.Now().Format("20060102"))
		orderId = orderId*100 + userId%100
	}
	return orderId
}
