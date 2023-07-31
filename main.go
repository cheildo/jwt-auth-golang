package main

import (
	"fmt"
	"net/http"

	"github.com/cheildo/jwt-auth-golang/login"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/login", login.LoginHandler).Methods("POST")
	router.HandleFunc("/protected", login.ProtectedHandler).Methods("GET")

	fmt.Println("Starting the server")
	err := http.ListenAndServe("localhost:4000", router)
	if err != nil {
		fmt.Println("Could not start the server", err)
	}
	fmt.Println("Server started. Listenning on port 4000")
}
