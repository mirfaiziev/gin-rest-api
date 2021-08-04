package server

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitHttpServer() {
	router := NewRouter()
	port := fmt.Sprintf(":%s", viper.Get("APP_PORT"))

	router.Run(port)
}
