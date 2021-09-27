package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmluci/KissatenService/internal/kissaten/models"
)

type KissatenRequest struct {
	Type     string `json:"type"`
	Username string `json:"usename,omitempty"`
	ItemName string `json:"item_name,omitempty"`
	Qty      string `json:"qty,omitempty"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

type FailedResponse struct {
	ErrorMessage string `json:"message"`
}

func (req *KissatenRequest) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(req)
}

func (sr *SuccessResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (fr *FailedResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(fr)
}

func BuyItemController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "BUY_ITEM" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func ReturnItemController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json");

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "RETURN_ITEM" {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func DropCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func PayCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func GetAllCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func GetCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func MakeNewCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
