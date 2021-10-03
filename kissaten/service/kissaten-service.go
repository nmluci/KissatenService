package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func BuyItem(order_id int, item_id int, sum int) error {
	return nil
}

func ReturnItem(order_id int, item_id int, sum int) error {
	return nil
}

func DropCart(order_id int) error {
	return nil
}

func PayCart(order_id int) error {
	return nil
}

func GetAllCart() error {
	return nil
}

func GetCart(uid int) error {
	return nil
}

func MakeNewCart(uname string) (int, error) {
	var (
		rawJson map[string]interface{}
		err     error
	)

	resp, err := http.Get(fmt.Sprintf("http://localhost:8081/api/user/u/%s", uname))
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()
	rawBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	err = json.Unmarshal(rawBody, &rawJson)
	fmt.Println(rawBody)
	if err != nil {
		fmt.Printf("Err REST: %s", err)
		return -1, err
	}
	// return uid, nil
	return -1, nil
}
