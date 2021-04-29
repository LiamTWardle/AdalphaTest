package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePageHandler)

	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	var h homePage = NewHomePage()
	message := h.Hello()
	_, err := fmt.Fprintf(w, "%s", message)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
