package main

import (
	"embed"
	"log"
	"net/http"
)

// assets holds our static web server content.
var assets embed.FS
var html []byte

func main() {
	fs := http.FileServer(http.FS(assets))
	http.Handle("/assets/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(html)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
