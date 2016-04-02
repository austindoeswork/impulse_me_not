package main

import (
	"net/http"
)

// HandleTest is a test handler
func HandleTest(w http.ResponseWriter, request *http.Request) {
	w.Write([]byte("hey, whats up"))
}
