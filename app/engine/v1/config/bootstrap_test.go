package config

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)



func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	a  := App{}
	a.Initialize()

	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func TestCreateUserSuccessScenario(t *testing.T) {
	payload := []byte(`{"id":"123","url":"tesst", "createdAt":"2019-10-29T20:34:41+05:30","updatedAt":"2019-10-29T20:34:41+05:30"}`)
	req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(payload))
	response := executeRequest(req)
	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody["responseType"] != "success" {
		t.Errorf("Expected success '. Got '%v'",  responseBody["responseType"])
	}

	if responseBody["message"] != "success" {
		t.Errorf("Expected success '. Got '%v'", responseBody["message"])
	}

	if responseBody["code"] != "200" {
		t.Errorf("Expected 200 '. Got '%v'",  responseBody["code"])
	}
}

func TestCreateUserErrorScenarioWithDateParseError(t *testing.T) {
	payload := []byte(`{"id":"1234","url":"tesst", "createdAt":"v","updatedAt":"2019-10-29T20:34:41+05:30"}`)
	req, _ := http.NewRequest("POST", "/api", bytes.NewBuffer(payload))
	response := executeRequest(req)
	var responseBody map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &responseBody)

	if responseBody["responseType"] != "error" {
		t.Errorf("Expected error '. Got '%v'",  responseBody["responseType"])
	}

	if 	!strings.Contains(responseBody["message"].(string),"cannot parse") {
		t.Errorf("Expected success '. Got '%v'", responseBody["message"])
	}

	if responseBody["code"] != "400" {
		t.Errorf("Expected 400 '. Got '%v'",  responseBody["code"])
	}
}
