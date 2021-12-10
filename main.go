package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Start")
		defer log.Println("End")

		longOperation(r.Context())
	})

	addr := ":8080"
	log.Printf("Listening on %q", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func longOperation(ctx context.Context) {
	v := 100

	for i := 0; i < v; i++ {
		select {
		case <-ctx.Done():
			log.Println("Abort")
			return
		default:
		}
		log.Printf("Operation: %v/%v", i, v)
		time.Sleep(time.Second * 1)
	}
}
