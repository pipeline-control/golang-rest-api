package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pipeline-control/golang-rest-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Server is up and running...")
	})

	router.HandleFunc("/create", routes.CreateUser).Methods("POST")
	router.HandleFunc("/getUsers", routes.GetUsers).Methods("GET")
	router.HandleFunc("/getUser", routes.GetUser).Methods("GET")
	router.HandleFunc("/update", routes.UpdateUser).Methods("PUT")
	router.HandleFunc("/delete", routes.DeleteUser).Methods("DELETE")

	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
