package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux" // Routes
)

const siteTitle = "Fy.to"





func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Page)
	r.HandleFunc("/contact", Contact )

	r.HandleFunc("/{page}", Page )
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.NotFoundHandler = http.HandlerFunc(NotFound)

	srv := &http.Server{
        Handler:      r,
        Addr:         "127.0.0.1:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    srv.ListenAndServe()
}

