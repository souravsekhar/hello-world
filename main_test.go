package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestHomeHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	http.HandlerFunc(homeHandler).
		ServeHTTP(rr, req)

	expected := string("Hello World!")
	assert.Equal(t, expected, rr.Body.String(), "Response body differs")
}
