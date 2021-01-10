package config

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type Config struct {
	DatabaseHost string
	DatabasePort string
	DatabaseUser string
	DatabasePass string
	DatabaseName string
}

func init() {
	viper.SetConfigFile("config.json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}
func NewConfig() *Config {

	conf := &Config{
		DatabaseHost: os.Getenv("DB_HOST"),
		DatabasePort: os.Getenv("DB_PORT"),
		DatabaseUser: os.Getenv("DB_USER"),
		DatabasePass: os.Getenv("DB_PASS"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
	return conf
}
func StarServer(router *mux.Router) {
	address := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	err := http.ListenAndServe(address, router)

	if err != nil {
		panic(err)
	}

}
