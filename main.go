package main

import (
	"jeloserver/api"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/feed", api.Feed)

	log.Println("Starting service at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
