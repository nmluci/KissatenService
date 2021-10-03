package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmluci/KissatenService/cafe/service"
)

type KissatenRequest struct {
	Type     string `json:"type"`
	Username string `json:"usename,omitempty"`
	ItemName string `json:"item_name,omitempty"`
	Qty      string `json:"qty,omitempty"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type FailedResponse struct {
	ErrorMessage string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
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

func BuyItem() http.HandlerFunc {
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

func ReturnItem() http.HandlerFunc {
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

func DropCart() http.HandlerFunc {
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

func PayCart() http.HandlerFunc {
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

func GetAllCart() http.HandlerFunc {
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

func GetCart() http.HandlerFunc {
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

func MakeNewCart() http.HandlerFunc {
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
			return
		}

		if orderId, err := service.MakeNewCart(order.Username); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &FailedResponse{
				ErrorMessage: "Failed to make new cart",
				Data:         err,
			}
			errResp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusOK)
			dataResp := &SuccessResponse{
				Message: "Cart has been made!",
				Data: map[string]int{
					"cartId": orderId,
				},
			}
			dataResp.ToJson(w)
		}
	}
}
