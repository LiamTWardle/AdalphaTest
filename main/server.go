package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	portfolio Portfolio
}

func NewServer(portfolio Portfolio) *server {
	s := new(server)
	s.portfolio = portfolio
	return s
}

func (s server) Listen() {
	http.HandleFunc("/", s.HomePageHandler)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func (s server) HomePageHandler(w http.ResponseWriter, r *http.Request) {
	instruction := s.portfolio.Withdraw(10)
	_, err := fmt.Fprintf(w, "%s", instruction)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
