package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/haunshila/go-ms/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(data.GetProducts())
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
	rw.Write(b)
}
