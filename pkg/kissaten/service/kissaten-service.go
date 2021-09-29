package service

import (
	"github.com/nmluci/KissatenService/internal/kissaten/models"
)

func BuyItem(km *models.KissatenModel, order_id int, item_id int, sum int) error {
	return nil
}

func ReturnItem(km *models.KissatenModel, order_id int, item_id int, sum int) error {
	return nil
}

func DropCart(km *models.KissatenModel, order_id int) error {
	return nil
}

func PayCart(km *models.KissatenModel, order_id int) error {
	return nil
}

func GetAllCart(km *models.KissatenModel) error {
	return nil
}

func GetCart(km *models.KissatenModel, uid int) error {
	return nil
}

func MakeNewCart(km *models.KissatenModel, uid int) (int, error) {
	return -1, nil
}
