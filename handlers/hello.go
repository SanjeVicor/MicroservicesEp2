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

func (self *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	self.l.Println("Hello World")

	//read the body
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error reading body", err)
		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

	//write the response
	fmt.Fprintf(rw, "Hello %s \n", d)
}
