package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// handler is an interface with one method - ServeHTTP(ResponseWriter, *Request)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Oops", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s", d)
}
