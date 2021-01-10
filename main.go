package main

import (
	"code-assignment/config"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.NewConfig()
	db := config.NewDatabase(conf)
	router := mux.NewRouter()

	config.NewRoutes(db, router)
	defer db.Close()
	config.StarServer(router)
}
