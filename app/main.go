package main

import "net/http"

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	}) // first endpoint returns Hello World
	router.Run(":8000")
}
