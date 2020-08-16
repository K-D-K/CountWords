package main

import (
	"WebScrapper/handler"
	"WebScrapper/handler/scrapper"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/scrape", handler.Executor(scrapper.Get)).Methods("GET")
	http.ListenAndServe(":8080", router)
}
