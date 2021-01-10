package utils

import (
	"encoding/json"
	"net/http"
)



func SendResponseSuccess(code int, message string, data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	result, _ := json.Marshal(data)
	w.Write(result)
}
