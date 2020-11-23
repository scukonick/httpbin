package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/scukonick/httpbin/internal/handlers/delay"
)

func main() {
	mux := &http.ServeMux{}

	mux.Handle("/delay", delay.NewHandler())

	server := &http.Server{
		ReadTimeout: 1 * time.Minute,
		Addr:        "127.0.0.1:9898",
		Handler:     mux,
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("startng to listen at %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("failed to start: %+v", err)
			return
		}
	}()

	wg.Wait()
}
