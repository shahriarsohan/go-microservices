package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// HandLeFunc is a convenient method. It takes my func ,create an http handler from it and add it to default servemux
	d, err := ioutil.ReadAll(r.Body)
	// log.Println(r.Method)
	log.Printf("%s", d)
	if err != nil {
		http.Error(w, "OOOPPP", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello %s", d)
}
