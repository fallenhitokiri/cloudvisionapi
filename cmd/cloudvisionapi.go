package main

import (
	"log"
	"net/http"

	"github.com/fallenhitokiri/cloudvisionapi"
)

func main() {
	http.HandleFunc("/", cloudvisionapi.Upload)
	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
