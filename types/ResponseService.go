package types

type ResponseService interface {
	ReturnExchange(pairs []string) map[string]float64
	ReturnExchangeByPair() map[string]float64
}
