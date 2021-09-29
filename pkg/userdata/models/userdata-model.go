package models

import (
	"database/sql"
	"errors"

	"github.com/mattn/go-sqlite3"
)

const (
	GET_ALL_USER         string = "SELECT id, username, credit FROM UserData"
	GET_USER_STMT        string = "SELECT id, username, credit FROM UserData where username = ?"
	INSERT_NEW_USER_STMT string = "INSERT INTO UserData(username) VALUES (?)"
)

type UserModel struct {
	DB *sql.DB
}

type UserData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Credit   int    `json:"credit"`
}

type Users []*UserData

func (um *UserModel) GetUserByName(uname string) (*UserData, error) {
	rows, err := um.DB.Query(GET_USER_STMT, uname)
	if err != nil {
		return nil, err
	}
	usr := &UserData{}

	defer rows.Close()
	rows.Next()
	rows.Scan(&usr.Id, &usr.Username, &usr.Credit)
	return usr, nil
}

func (um *UserModel) GetAllUser() (Users, error) {
	var users Users

	rows, err := um.DB.Query(GET_ALL_USER)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp UserData
		rows.Scan(&temp.Id, &temp.Username, &temp.Credit)
		users = append(users, &temp)
	}
	return users, nil
}

func (um *UserModel) RegisterNewUser(uname string) error {
	if _, err := um.DB.Exec(INSERT_NEW_USER_STMT, uname); err != nil {
		if driverErr, ok := err.(sqlite3.Error); ok {
			if driverErr.ExtendedCode == 2067 {
				return errors.New("existing username detected")
			}
		}
		return err
	} else {
		return nil
	}
}
