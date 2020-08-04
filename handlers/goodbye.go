package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (self *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	self.l.Println("Goodbye World")
	rw.Write([]byte("Byeee! \n"))
}
