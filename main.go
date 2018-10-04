/*
	gasquez.net - website
	~~~~~~~~~~~~~~~~~~~~~
	:license: BSD, see LICENSE for more details.
*/

package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const siteTitle = "Fy.to"
const debug = false

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Page)
	r.HandleFunc("/contact", Contact )
	r.HandleFunc("/projects", Projects )

	r.HandleFunc("/{page}", Page )
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	r.NotFoundHandler = http.HandlerFunc(NotFound)

	srv := &http.Server{
        Handler:      r,
        Addr:         "0.0.0.0:8000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    srv.ListenAndServe()
}

