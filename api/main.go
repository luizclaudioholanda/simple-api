package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.LoadConfig()

	fmt.Println("Executando API!")

	route := router.Create()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), route))
}
