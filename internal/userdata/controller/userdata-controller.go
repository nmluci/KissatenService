package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/nmluci/KissatenService/internal/userdata/models"
	"github.com/nmluci/KissatenService/internal/userdata/service"
	"golang.org/x/text/message"
)

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error,omitempty"`
	ErrCode      int    `json:"code,omitempty"`
}

func (sr *SuccessResponse) ToJson(w *io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (er *ErrorResponse) ToJson(w *io.Writer) {
	json.NewEncoder(w).Encode(er)
}

func GetAllUserdataController(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content=Type", "application/json")

		usr, err := service.GetAllUser(um)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &ErrorResponse{
				message: err.Error()
			}
			errResp.ToJson(w)
		} else {
			resp := &SuccessResponse{
				Data: usr,
			}
			resp.ToJson(w)
		}
	}
}

func GetUserDataController(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func RegisterNewUser(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
