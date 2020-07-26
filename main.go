package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"paste-corral/pastebin"
	"paste-corral/rest"
)

func main() {

	// -- --------------------------------------
	// Run web crawler in concurrent goroutine:
	go pastebin.Crawl()

	// -- --------------------------------------
	// Get env info:
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	p("Paste Corral", version(), "started at port", port)

	// -- --------------------------------------
	// Setup server mux:
	mux := http.NewServeMux()

	// For file server purposes:
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	// The main handler:
	mux.HandleFunc("/", rest.HandleRequest)

	// starting up the server
	server := &http.Server{
		Addr:           ":" + port,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
