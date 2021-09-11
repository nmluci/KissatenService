package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/internal/userdata/models"
	"github.com/nmluci/KissatenService/internal/userdata/service"
)

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // for future-proof
}

type SuccessResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	ErrorMessage string `json:"error,omitempty"`
	ErrCode      int    `json:"code,omitempty"`
}

func (usr *NewUser) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(usr)
}

func (sr *SuccessResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(sr)
}

func (er *ErrorResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(er)
}

func GetAllUserdataController(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content=Type", "application/json")

		usr, err := service.GetAllUser(um)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &ErrorResponse{
				ErrorMessage: err.Error(),
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
		w.Header().Add("Content-Type", "application/json")

		usernameQuery, exist := mux.Vars(r)["username"]
		if !exist {
			w.WriteHeader(http.StatusBadRequest)
			err := &ErrorResponse{
				ErrorMessage: "No username!",
			}
			err.ToJson(w)
			return
		}

		usr, err := service.GetUserByName(um, usernameQuery)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			errResp := &ErrorResponse{
				ErrorMessage: err.Error(),
			}
			errResp.ToJson(w)
		} else if usr.Username == "" {
			w.WriteHeader(http.StatusNotFound)
			resp := &ErrorResponse{
				ErrorMessage: "User not found!",
			}
			resp.ToJson(w)
		} else {
			resp := &SuccessResponse{
				Data: usr,
			}
			resp.ToJson(w)
		}
	}
}

func RegisterNewUser(um *models.UserModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		newUser := &NewUser{}
		newUser.FromJson(r.Body)

		if err := service.RegisterNewUser(um, newUser.Username); err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			errResp := &ErrorResponse{
				ErrorMessage: err.Error(),
			}
			errResp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}
