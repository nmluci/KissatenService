package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/nmluci/KissatenService/internal/inventory/model"
	"github.com/nmluci/KissatenService/internal/inventory/service"
)

type IncomingData struct {
	UpdateType *int    `json:"updateType,omitempty"` // 0  all, 1 price, 2 stocks
	ItemId     *int    `json:"itemId,omitempty"`
	Name       *string `json:"name,omitempty"`
	Price      *int    `json:"price,omitempty"`
	Stocks     *int    `json:"stocks,omitempty"`
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
			log.Println(err)
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

		if itemId, exists := mux.Vars(r)["itemId"]; exists {
			itemId, err := strconv.Atoi(itemId)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errResp := &ErrorResponse{Message: err.Error(), ErrCode: 1}
				errResp.ToJson(w)
				return
			}

			if itm, err := service.GetItemById(im, itemId); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				errResp := &ErrorResponse{Message: err.Error(), ErrCode: 1}
				errResp.ToJson(w)
			} else {
				resp := &SuccessResponse{Data: itm}
				resp.ToJson(w)
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("itemId isn't valid")
			errResp := &ErrorResponse{Message: "itemId isn't valid"}
			errResp.ToJson(w)
		}

	}
}

func RemoveItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if itemId, exists := mux.Vars(r)["itemId"]; exists {
			itemId, err := strconv.Atoi(itemId)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				errResp := &ErrorResponse{Message: err.Error(), ErrCode: 1}
				errResp.ToJson(w)
				return
			}

			if err := service.RemoveItem(im, itemId); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				errResp := &ErrorResponse{Message: err.Error()}
				errResp.ToJson(w)
			} else {
				resp := &SuccessResponse{Message: "Item has been removed"}
				resp.ToJson(w)
			}
		} else {
			log.Println("itemId isn't valid")
			w.WriteHeader(http.StatusBadRequest)
			errResp := &ErrorResponse{Message: "itemId isn't valid"}
			errResp.ToJson(w)
		}
	}
}

func UpdateItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		itmId, exists := mux.Vars(r)["itemId"]
		if !exists {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("itemId isn't valid")
			errResp := &ErrorResponse{Message: "itemId isn't valid"}
			errResp.ToJson(w)
			return
		}

		itemId, err := strconv.Atoi(itmId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			errResp := &ErrorResponse{Message: err.Error()}
			errResp.ToJson(w)
			return
		}

		reqBody := &IncomingData{}
		reqBody.FromJson(r.Body)

		if req := reqBody.UpdateType; req != nil {
			var err error
			switch *req {
			case 0:
				newitem := &models.Item{
					Id:    *reqBody.ItemId,
					Name:  *reqBody.Name,
					Price: *reqBody.Price,
					Stock: *reqBody.Stocks,
				}

				err = service.UpdateItem(im, itemId, newitem)
			case 1:
				err = service.UpdateItemPrice(im, itemId, *reqBody.Price)
			case 2:
				err = service.UpdateItemStocks(im, itemId, *reqBody.Stocks)
			default:
				err = errors.New("invalid update code")
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				errRes := &ErrorResponse{
					Message: err.Error(),
				}
				errRes.ToJson(w)
				return
			}

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				errRes := &ErrorResponse{
					Message: err.Error(),
				}
				errRes.ToJson(w)
			} else {
				resp := &SuccessResponse{
					Message: "Success",
				}
				resp.ToJson(w)
			}

		} else {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("No update type specified")
			errResp := &ErrorResponse{Message: "no update type specified"}
			errResp.ToJson(w)
		}
	}
}

func InsertItemController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		reqBody := &IncomingData{}
		reqBody.FromJson(r.Body)
		if reqBody == nil || reqBody.Name == nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println("no data given")
			errResp := &ErrorResponse{
				Message: "seriously? _-",
				ErrCode: 0,
			}
			errResp.ToJson(w)
			return
		}

		if reqBody.Price == nil {
			*reqBody.Price = 0
		}

		if reqBody.Stocks == nil {
			*reqBody.Price = 0
		}

		if itemId, err := service.InsertItem(im, *reqBody.Name, *reqBody.Price, *reqBody.Stocks); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
			errResp := &ErrorResponse{
				Message: err.Error(),
			}
			errResp.ToJson(w)
			return
		} else {
			resp := &SuccessResponse{
				Message: fmt.Sprintf("item registered with id: %d", *itemId),
			}
			resp.ToJson(w)
		}

	}
}

func GetItemByQueryController(im *models.InventoryModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		if itemName, exists := mux.Vars(r)["itemName"]; exists {
			itm, err := service.GetItemByName(im, itemName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Println(err)
				errResp := &ErrorResponse{Message: err.Error()}
				errResp.ToJson(w)
			}
			resp := &SuccessResponse{Data: itm}
			resp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			err := &ErrorResponse{
				Message: "item not found in the database",
				ErrCode: 404,
			}
			err.ToJson(w)
		}
	}
}
