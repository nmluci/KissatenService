package service

import (
	"github.com/nmluci/KissatenService/internal/userdata/models"
)

func GetUserByName(um *models.UserModel, uname string) (*models.UserData, error) {
	if data, err := um.GetUserByName(uname); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func GetAllUser(um *models.UserModel) (models.Users, error) {
	if data, err := um.GetAllUser(); err != nil {
		return nil, err
	} else {
		return data, nil
	}
}

func RegisterNewUser(um *models.UserModel, uname string) error {
	if err := um.RegisterNewUser(uname); err != nil {
		return err
	} else {
		return nil
	}
}
