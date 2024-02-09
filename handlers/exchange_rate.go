package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/iriojose/cada_test/clients"
	"github.com/iriojose/cada_test/types"
	"github.com/iriojose/cada_test/utils"
)

func ExchangeRateHandler(w http.ResponseWriter, r *http.Request) error {
	var exchangeRateRequest types.ExchangeRateRequest
	json.NewDecoder(r.Body).Decode(&exchangeRateRequest)

	validation := ValidateRequest(exchangeRateRequest)
	if validation != nil {
		return validation
	}

	pairs := strings.Split(exchangeRateRequest.CurrencyPair, "-")

	response := make(chan types.ResponseService)
	errorResponse := make(chan error)

	go clients.ServiceA(pairs, response, errorResponse)
	go clients.ServiceB(pairs, response, errorResponse)

	responseData := <-response
	err := <-errorResponse

	if err != nil {
		return err
	}

	var exchangeRateResponse types.ExchangeRateResponse = responseData.ReturnExchange(pairs)
	return utils.WriteJson(w, http.StatusOK, exchangeRateResponse)
}

func ValidateRequest(exchangeRateRequest types.ExchangeRateRequest) error {
	if exchangeRateRequest.CurrencyPair == "" {
		return fmt.Errorf("Currency pair is required")
	}

	//regex format "XXX-XXX"
	validFormat := regexp.MustCompile(`^[a-zA-Z]+-[a-zA-Z]+$`)

	if !validFormat.MatchString(exchangeRateRequest.CurrencyPair) {
		return fmt.Errorf("Invalid currency pair format. Should be in the format 'XXX-XXX'")
	}
	return nil
}
