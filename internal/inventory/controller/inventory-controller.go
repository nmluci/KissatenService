package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (sr *SuccessResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (er *ErrorResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(er)
}

func GetAllItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func GetItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func RemoveItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func UpdateItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func InsertItemController() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

