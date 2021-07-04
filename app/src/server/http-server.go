package server

func InitHttpServer() {
	router := NewRouter()
	router.Run(":8081")
}
