package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApiPost(t *testing.T) {

	var jsonStr = []byte(`{"id":4,"url":"test","created_at":"","updated_at":"","active":1}`)

	req, err := http.NewRequest("POST", "/api", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ApiPost)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	if rr.Body.String() != "success" {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), "success")
	}

}
