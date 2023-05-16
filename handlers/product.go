package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	json.Marshal
}
