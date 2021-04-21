package main

import (
	"log"
	"net/http"
	"os"
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

	s.ListenAndServe()

}
