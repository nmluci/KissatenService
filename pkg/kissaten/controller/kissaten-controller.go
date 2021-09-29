package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmluci/KissatenService/internal/kissaten/models"
	"github.com/nmluci/KissatenService/internal/kissaten/service"
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
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "RETURN_ITEM" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func DropCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "DROP_CART" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func PayCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "PAY_NOW" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func GetAllCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "GET_ALL" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func GetCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}
	}
}

func MakeNewCartController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		order := &KissatenRequest{}
		order.FromJson(r.Body)

		if order.Type != "NEW_CART" {
			w.WriteHeader(http.StatusBadRequest)
			errResp := &FailedResponse{
				ErrorMessage: "Order type invalid",
			}
			errResp.ToJson(w)
		}

		var orderId = service.MakeNewCart(km, order.Username)
	}
}
