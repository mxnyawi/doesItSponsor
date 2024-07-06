package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mxnyawi/doesItSponsor/internal/handler"
)

func TestDoesItSponseHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/doesitsponse/http://example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handler.DoesItSponsorHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := `{"message":"The title does not contain the word 'example'."}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
