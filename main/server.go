package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct {
	homePage HomePage
}

func NewServer(homePage HomePage) *server {
	s := new(server)
	s.homePage = homePage
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
	message := s.homePage.Hello()
	_, err := fmt.Fprintf(w, "%s", message)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
