package middleware

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type FailedResponse struct {
	ErrorMessage string `json:"message"`
}

func (fr *FailedResponse) ToJson(w io.Writer) {
	json.NewEncoder(w).Encode(fr)
}
func NoFunc() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNotImplemented)

			msg := FailedResponse{
				ErrorMessage: "What art ye doin",
			}
			msg.ToJson(w)
		})
	}
}
