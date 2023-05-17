package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("error in ListenAndServe(). err=%v", err)
	}
}
