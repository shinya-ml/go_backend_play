package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("pong")
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatalf("error in ListenAndServe(). err=%v", err)
	}
}
