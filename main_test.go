package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/?name=test", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(testHandler)

	handler.ServeHTTP(rr, req)

	expectStatus := http.StatusOK
	if status := rr.Code; status != expectStatus {
		t.Errorf(
			"handler returned wrong status code: got %v want %v",
			status, expectStatus,
		)
	}

	expected := "Hello test!"
	if rr.Body.String() != expected {
		t.Errorf(
			"handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected,
		)
	}
}
