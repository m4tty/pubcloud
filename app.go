package main

import (
	"github.com/m4tty/pubcloud/web/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/pubs", handlers.GetPubsHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/dist/"))))
	http.ListenAndServe(":8080", nil)
}
