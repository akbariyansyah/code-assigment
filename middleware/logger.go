package middleware

import (
	"log"
	"net/http"
)

func ActivityLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Printf("%v - %v \t from %v ", req.Method,req.RequestURI,req.RemoteAddr)
		next.ServeHTTP(resp, req)
	})
}
