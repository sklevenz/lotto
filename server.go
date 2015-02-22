package main

import (
	"log"
	"net/http"
	"time"
)

func createHandler() http.Handler {
	mux := http.NewServeMux()

	// handle static assets
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/random", random)

	return mux
}

func startServer() {
	log.Println("-- lotto app", config.Version, "started at", config.Address, "--------------------------")
	// starting up the server
	server := &http.Server{
		Addr:           config.Address,
		Handler:        createHandler(),
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()

}
