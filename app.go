package main

import (
	"github.com/m4tty/pubcloud/web/handlers"
	"net/http"
)

//var router = new(mux.Router)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/pubs", handlers.GetPubsHandler).Name("pubs-getall").Methods("GET")
	// http.Handle("/", r)
	http.HandleFunc("/pubs", handlers.GetPubsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/dist/"))))
	http.ListenAndServe(":8080", nil)
}
