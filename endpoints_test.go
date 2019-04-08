package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler).Methods("GET")
	return router
}

func TestRootHandler(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code, "Status 200: OK expected.")
	assert.Equal(t, "Hello World!\n", response.Body.String(), "Expect body of \"Hello World!\"")
}
