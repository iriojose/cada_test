package types

import "fmt"

type ServiceAResponse struct {
	Base          string             `json:"base"`
	LastUpdated   int64              `json:"last_updated"`
	ExchangeRates map[string]float64 `json:"exchange_rates"`
}

func (s ServiceAResponse) ReturnExchange(pairs []string) map[string]float64 {
	code := fmt.Sprintf("%s-%s", pairs[0], pairs[1])
	return map[string]float64{code: s.ExchangeRates[pairs[1]]}
}

func (s ServiceAResponse) ReturnExchangeByPair() map[string]float64 {
	return s.ExchangeRates
}
