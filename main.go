package main

import (
	"net/http"

	"github.com/souvikhaldar/go-streamer/downloader"
	"github.com/souvikhaldar/go-streamer/uploader"
	"github.com/souvikhaldar/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", uploader.Upload).Methods("POST")
	router.HandleFunc("/{ID}", downloader.Download).Methods("GET")
	http.ListenAndServe(":8192", router)
}
