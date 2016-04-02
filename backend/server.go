package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// startServer starts running the server and blocks until the
// program exists
func startServer() {
	fmt.Println("start server")
	http.Handle("/", newHandler())
}

func newHandler() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/test", HandleTest).Methods("GET")

	return router
}
