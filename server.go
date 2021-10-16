package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Server is up and running...")
	})

	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
