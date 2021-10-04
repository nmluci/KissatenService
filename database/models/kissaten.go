package models

type KissatenModel struct {
	OrderId  int `json:"order_id"`
	MemberId int `json:"member_id"`
	ItemId   int `json:"item_id"`
	Sum      int `json:"sum"`
}

type KissatenDB struct {
	Size int             `json:"size"`
	Data []*KissatenModel `json:"data"`
}
