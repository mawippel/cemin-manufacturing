package main

import (
	"controleConfeccao/internal/api"
	"controleConfeccao/internal/db"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	router := api.SetupRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
