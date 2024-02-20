package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpGET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from server"))
	}))
	defer server.Close()

	response, err := HttpGET(server.URL)
	if err != nil {
		t.Fatalf("Error en la llamada HTTP GET: %v", err)
	}

	expectedResponse := []byte("Response from server")
	if !bytes.Equal(response, expectedResponse) {
		t.Errorf("Wrong answer. expected %s but got %s", string(expectedResponse), string(response))
	}
}

func TestHttpPOST(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("A POST method was expected, but got %s", r.Method)
		}

		body, _ := ioutil.ReadAll(r.Body)
		expectedBody := []byte("Request body")
		if !bytes.Equal(body, expectedBody) {
			t.Errorf("wrong boy. expected %s but got %s", string(expectedBody), string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Response from server"))
	}))
	defer server.Close()

	response, err := HttpPOST(server.URL, "text/plain", []byte("Request body"))
	if err != nil {
		t.Fatalf("Error en la llamada HTTP POST: %v", err)
	}

	expectedResponse := []byte("Response from server")
	if !bytes.Equal(response, expectedResponse) {
		t.Errorf("Wrong answer. expected %s but got %s", string(expectedResponse), string(response))
	}
}
