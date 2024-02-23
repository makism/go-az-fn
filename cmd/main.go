package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("api/service", handlerRoot)
	router.GET("api/service/ping", handlerPing)

	portInfo := GetApiPort()
	router.Run(portInfo)
	log.Println("API is up & running - " + portInfo)
}

func GetApiPort() string {
	port := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		port = ":" + val
	}
	return port
}

func handlerRoot(c *gin.Context) {
	log.Println("Invoke ROOT")
	c.JSON(http.StatusOK, gin.H{
		"message": "Okay",
	})
}

func handlerPing(c *gin.Context) {
	log.Println("Invoke PING")
	c.JSON(http.StatusOK, gin.H{
		"message": "Pong!",
	})
}