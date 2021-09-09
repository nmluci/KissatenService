package controller

import (
	"encoding/json"
	"io"
	"net/http"

	models "github.com/nmluci/KissatenService/internal/inventory/model"
	"github.com/nmluci/KissatenService/internal/inventory/service"
)

type IncomingData struct {
	ItemId *int    `json:"itemId,omitempty"`
	Name   *string `json:"name,omitempty"`
	Price  *int    `json:"price,omitempty"`
	Stocks *int    `json:"stocks,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	ErrCode int    `json:"error_code"`
	Message string `json:"message"`
}

func (id *IncomingData) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(id)
}

func (sr *SuccessResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (er *ErrorResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(er)
}

func GetAllItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		itm, err := service.GetAllItem(im)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errRes := &ErrorResponse{
				ErrCode: 101,
				Message: err.Error(),
			}
			errRes.ToJson(w)
		}

		res := &SuccessResponse{
			Data: itm,
		}
		res.ToJson(w)
	}
}

func GetItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		ReqBody := &IncomingData{}
		ReqBody.FromJson(r.Body)
	}
}

func RemoveItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func UpdateItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func InsertItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
