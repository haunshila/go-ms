package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/haunshila/go-ms/handlers"
)

var bindAddress = ":9090"

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	pd := handlers.NewProducts(l)

	// create a server mux and assign handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()

	getRouter.HandleFunc("/", pd.GetProducts)
	// sm.Handle("/products", pd)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", pd.UpdateProducts)
	putRouter.Use(pd.MiddlewareValidateProduct)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", pd.AddProducts)
	postRouter.Use(pd.MiddlewareValidateProduct)

	s := &http.Server{
		Addr:         bindAddress,
		Handler:      sm,
		IdleTimeout:  60 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")
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
