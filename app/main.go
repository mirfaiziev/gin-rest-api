package main

import (
	"rest-api-app/src/server"
)

func main() {
	server.InitConfig()
	server.InitDatabase()
	//server.InitHttpServer()
}
