package service

import (
	"errors"

	"github.com/nmluci/KissatenService/database/models"
)

func IsServiceExists(serviceName string, services *[]string) bool {
	for _, data := range *services {
		if data == serviceName {
			return true
		}
	}
	return false
}

func stringToInterface(origin []string) []interface{} {
	var dest []interface{}

	for _, data := range origin {
		dest = append(dest, data)
	}

	return dest
}

func GetItem(db *models.DatabaseModel, query string, param []string, serviceName string) (interface{}, error) {
	if serviceName == "kissaten" {
		resp, err := KissatenGetItem(db, query, stringToInterface(param))
		if err != nil {
			return nil, err
		} else {
			return resp, nil
		}
	} else {
		return nil, nil
	}
}

func PostItem(db *models.DatabaseModel, query string, param []string, serviceName string) error {
	if serviceName == "kissaten" {
		err := KissatenPostItem(db, query, stringToInterface(param))
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func RegisterService(serviceName string, services *[]string) error {
	if !IsServiceExists(serviceName, services) {
		*services = append(*services, serviceName)
		return nil
	} else {
		return errors.New("services already registered")
	}
}
