package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start")
		defer log.Println("End")

		longOperation()
	})

	addr := ":8080"
	log.Printf("Application is listening on %q", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func longOperation() {
	v := 100
	for i := 0; i < v; i++ {
		log.Printf("Operation: %v/%v", i, v)
		time.Sleep(time.Second * 1)
	}
}
