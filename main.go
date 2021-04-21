package main

import (
	"log"
	"net/http"
	"os"

	"github.com/haunshila/go-ms/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	hl := handlers.NewHello(l)
	gb := handlers.NewGoodBye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hl)
	sm.Handle("/bye", gb)

	http.ListenAndServe(":9090", sm)

}
