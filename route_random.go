package main

import (
	"log"
	"net/http"
)

func random(writer http.ResponseWriter, request *http.Request) {
	log.Println("/random")
	http.Redirect(writer, request, "/", http.StatusFound)
}
