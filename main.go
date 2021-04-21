package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/haunshila/go-ms/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hl := handlers.NewHello(l)
	gb := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hl)
	sm.Handle("/bye", gb)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Received terminate, graceful shutdown", sig)

	// Graceful shutdown
	tc, _ := context.WithDeadline(context.Background(), time.Now().Add(30*time.Second))
	s.Shutdown(tc)

}
