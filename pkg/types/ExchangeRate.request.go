package types

type ExchangeRateRequest struct {
	CurrencyPair string `json:"currency-pair" validate:"required"`
}
