package scrapper

import (
	"WebScrapper/controller/scrapper"
	"WebScrapper/handler"
	"fmt"
	"net/http"
)

// Get handler for Scrapper
func Get(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query["url"][0]
	fmt.Println(url)
	byteArr := scrapper.Scrape(url)

	handler.RespondwithJSON(w, http.StatusOK, byteArr)
}
