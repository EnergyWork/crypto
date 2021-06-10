package main

import (
	"api-crypto/controllers"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func HandleRequests() {
	router := gin.Default()
	router.GET("/api/crypto", controllers.CheckAccess)
	router.POST("/api/crypto", controllers.Crypto)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err := router.Run(":" + port)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main() {
	HandleRequests()
	//debug_decrypt()
}
