package service

import (
	"github.com/nmluci/KissatenService/database/models"
)

func KissatenGetItem(db models.DatabaseModel, query string, param []interface{}) (*models.KissatenDB, error) {
	data := &models.KissatenDB{}
	rows, err := db.DB.Query(query, param...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var temp models.KissatenModel
		rows.Scan(&temp.OrderId, &temp.MemberId, &temp.ItemId, &temp.Sum)
		data.Data = append(data.Data, &temp)
		data.Size++
	}
	return data, nil
}

func KissatenPostItem(db models.DatabaseModel, query string, param []interface{}) error {
	if _, err := db.DB.Exec(query, param...); err != nil {
		return err
	}
	return nil
}
