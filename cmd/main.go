package main

import (
	"azfn/pkg/routes"

	"fmt"
	"log"
	"net/http"
	"os"
)

func GetApiPort() string {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func main() {
	customHandlerPort := GetApiPort()

	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HandleRoot)
	mux.HandleFunc("/service-blobtrigger", routes.HandleBlobTrigger)

	fmt.Println("Go server Listening...on FUNCTIONS_CUSTOMHANDLER_PORT:", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux))
}