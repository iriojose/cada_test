package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/iriojose/cada_test/handlers"
	"github.com/iriojose/cada_test/utils"
)

// test exchange rate handler
/* func TestExchangeHandler(t *testing.T) {
	initializers.LoadEnv()
	server := httptest.NewServer(http.HandlerFunc(utils.MakeHttpHandlerFunc(handlers.ExchangeRateHandler)))

	requestBody := `{"currency-pair": "USD-EUR"}`
	expected := `{"USD-EUR": 0.929541}`
	res, err := http.Post(server.URL+"/exchange-rate", "application/json", bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
} */

// test exisiting currency pair
func TestCurrencyPair(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(utils.MakeHttpHandlerFunc(handlers.ExchangeRateHandler)))

	requestBody := `{}`
	expected := `{"Error":"Currency pair is required"}`
	res, err := http.Post(server.URL+"/exchange-rate", "application/json", bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}

// test format currency pair
func TestCurrencyFormat(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(utils.MakeHttpHandlerFunc(handlers.ExchangeRateHandler)))

	requestBody := `{"currency-pair": "USDEUR"}`
	expected := `{"Error":"Invalid currency pair format. Should be in the format 'XXX-XXX'"}`
	res, err := http.Post(server.URL+"/exchange-rate", "application/json", bytes.NewBuffer([]byte(requestBody)))

	if err != nil {
		t.Error(err)
	}

	defer res.Body.Close()

	if status := res.StatusCode; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		t.Error(err)
	}

	if strings.TrimSpace(string(body)) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}
