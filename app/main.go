package main

import (
	"rest-api-app/src/server"
)

func main() {
	server.InitConfig()
	//db := server.GetDB()

	// var result string
	// conn.Raw("select 1+1 as res").Scan(&result)

	// fmt.Println(result)

	server.InitHttpServer()
}
