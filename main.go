package main

import (
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/feed", Feed)

	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "5000"
	}

	log.Println("Starting service at port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
