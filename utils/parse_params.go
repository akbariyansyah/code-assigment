package utils

import (
	"github.com/gorilla/mux"
	"net/http"
)

// ParseParams -> parse incoming parameter inside request url
func ParseParams(key string, r *http.Request) string {
	param := mux.Vars(r)
	return param[key]
}
