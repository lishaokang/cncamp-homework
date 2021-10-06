package main

import (
	"io"
	"log"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("User-Agent", r.Header.Get("User-Agent"))
	io.WriteString(w, r.Header.Get("User-Agent"))
}

func healthz(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
