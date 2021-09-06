package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type FailedResponse struct {
	ErrorMessage string `json:"message"`
}

func (sr *SuccessResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (fr *FailedResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(fr)
}

func BuyItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func ReturnItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func DropCartController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func PayCartController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func GetAllCartController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func GetCartController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func MakeNewCartController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
