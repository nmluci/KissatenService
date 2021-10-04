package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
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

func (res *DatabaseResponse) FromJson(r io.Reader) {
	json.NewDecoder(r).Decode(res)
}

func (req *DatabaseRequest) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(req)
}

func RegisterService() error {
	var jsonData = new(bytes.Buffer)
	msg := &DatabaseRequest{
		Module:  "REGISTER_MODULE",
		Service: "kissaten",
	}
	msg.ToJson(jsonData)

	resp, err := http.Post("http://localhost:8085/api/database", "application/json", jsonData)
	if err != nil {
		return err
	}

	msgResp := &DatabaseResponse{}

	msgResp.FromJson(resp.Body)
	defer resp.Body.Close()
	log.Printf("Messenger: %s\n", msgResp.Status)
	if msgResp.Status == "SUCCESS" {
		return nil
	} else {
		data := msgResp.Data.(map[string]interface{})
		return errors.New(data["message"].(string))
	}
}
