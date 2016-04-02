package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
)

func getHTTPResponse(url string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", url, nil)
	newHandler().ServeHTTP(recorder, req)

	return recorder
}

func postHTTPResponse(url, body string) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", url, strings.NewReader(body))
	newHandler().ServeHTTP(recorder, req)

	return recorder
}
