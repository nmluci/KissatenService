package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmluci/KissatenService/internal/kissaten/models"
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

func BuyItemController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}

func ReturnItemController(km *models.KissatenModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
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
