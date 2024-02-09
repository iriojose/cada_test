package types

type GetPersonsResponse struct {
	Persons           []Person       `json:"persons"`
	ByAscendingSalary []Person       `json:"byAscendingSalary"`
	ByDecendingSalary []Person       `json:"byDecendingSalary"`
	CurrencyGroups    CurrencyGroups `json:"currencyGroups"`
	ByFilter          []Person       `json:"byFilter"`
}
