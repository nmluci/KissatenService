package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nmluci/KissatenService/database/models"
	"github.com/nmluci/KissatenService/database/service"
)

type DatabaseResponse struct {
	Module string      `json:"module"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type DatabaseRequest struct {
	Module  string   `json:"module,omitempty"`  // OP name
	Service string   `json:"service,omitempty"` // Service name
	Query   string   `json:"query,omitempty"`   // Query
	Params  []string `json:"params,omitempty"`  // Params
}

type SuccessResponse struct {
	Type    string      `json:"type"`
	Message interface{} `json:"message,omitempty"`
}

type FailedResponse struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"errorCode"`
}

func (res *DatabaseResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(res)
}

func (req *DatabaseRequest) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(req)
}

func GetItem(db *models.DatabaseModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		queryReq := &DatabaseRequest{}
		queryReq.FromJson(r.Body)
		serviceName := mux.Vars(r)["service"]

		if serviceName == "" || !service.IsServiceExists(serviceName, &db.Services) {
			w.WriteHeader(http.StatusBadRequest)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Status: "Rejected",
				Data: &FailedResponse{
					ErrorMessage: "Service not registered or invalid!",
					ErrorCode:    "DB-403",
				},
			}
			resp.ToJson(w)
			return
		}

		data, err := service.GetItem(db, queryReq.Query, queryReq.Params, serviceName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Status: "ERROR",
				Data: &FailedResponse{
					ErrorMessage: err.Error(),
					ErrorCode:    "DB-10",
				},
			}
			resp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusOK)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Data: &SuccessResponse{
					Type:    "GetItem",
					Message: data,
				},
			}
			resp.ToJson(w)
		}
	}
}

func PostItem(db *models.DatabaseModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		queryReq := &DatabaseRequest{}
		queryReq.FromJson(r.Body)
		serviceName := mux.Vars(r)["service"]

		if serviceName == "" || !service.IsServiceExists(serviceName, &db.Services) {
			w.WriteHeader(http.StatusBadRequest)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Status: "Rejected",
				Data: &FailedResponse{
					ErrorMessage: "Service not registered or invalid!",
					ErrorCode:    "DB-403",
				},
			}
			resp.ToJson(w)
			return
		}

		err := service.PostItem(db, queryReq.Query, queryReq.Params, serviceName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Status: "ERROR",
				Data: &FailedResponse{
					ErrorMessage: err.Error(),
					ErrorCode:    "DB-10",
				},
			}
			resp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusOK)
			resp := &DatabaseResponse{
				Module: queryReq.Module,
				Status: "SUCCESS",
				Data: &SuccessResponse{
					Type: "GetItem",
				},
			}
			resp.ToJson(w)
		}
	}
}

func RegisterService(db *models.DatabaseModel) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		registerReq := &DatabaseRequest{}
		registerReq.FromJson(r.Body)

		if registerReq.Module != "REGISTER_MODULE" {
			w.WriteHeader(http.StatusBadRequest)
			resp := &DatabaseResponse{
				Module: "NOT_FOUND",
				Status: "ERROR",
				Data: &FailedResponse{
					ErrorMessage: "Module not specified or invalid!",
					ErrorCode:    "DB-404",
				},
			}
			resp.ToJson(w)
			return
		}

		if err := service.RegisterService(registerReq.Service, &db.Services); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			resp := &DatabaseResponse{
				Module: "REGISTER_SERVICE",
				Status: "FAIELD",
				Data: &FailedResponse{
					ErrorMessage: err.Error(),
					ErrorCode:    "DB-500",
				},
			}
			resp.ToJson(w)
		} else {
			w.WriteHeader(http.StatusOK)
			resp := &DatabaseResponse{
				Module: "REGISTER_SERVICE",
				Status: "SUCCESS",
				Data: &SuccessResponse{
					Type:    "REGISTER",
					Message: db.Services,
				},
			}
			resp.ToJson(w)
		}
	}
}
