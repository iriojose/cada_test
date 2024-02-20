package types

import "fmt"

type ServiceBResponse struct {
	Result             string  `json:"result"`
	Documentation      string  `json:"documentation"`
	TermsOfUse         string  `json:"terms_of_use"`
	TimeLastUpdateUnix int     `json:"time_last_update_unix"`
	TimeLastUpdateUtc  string  `json:"time_last_update_utc"`
	TimeNextUpdateUnix int     `json:"time_next_update_unix"`
	TimeNextUpdateUtc  string  `json:"time_next_update_utc"`
	BaseCode           string  `json:"base_code"`
	TargetCode         string  `json:"target_code"`
	ConversionRate     float64 `json:"conversion_rate"`
}

func (s ServiceBResponse) ReturnExchange(pairs []string) map[string]float64 {
	code := fmt.Sprintf("%s-%s", pairs[0], pairs[1])
	return map[string]float64{code: s.ConversionRate}
}

func (s ServiceBResponse) ReturnExchangeByPair() map[string]float64 {
	return map[string]float64{s.TargetCode: s.ConversionRate}
}
