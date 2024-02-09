package types

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/iriojose/cada_test/utils"
)

// model
type Person struct {
	Id         int    `json:"id"`
	PersonName string `json:"personName"`
	Salary     Salary `json:"salary"`
}

// interfaces
type SortByAscending []Person
type SortByDecending []Person

type CurrencyGroups map[string][]Person
type GroupByCurrency []Person

type FilterBySalary []Person

// functions
func (s SortByAscending) Len() int           { return len(s) }
func (s SortByAscending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByAscending) Less(i, j int) bool { return s[i].Salary.Value < s[j].Salary.Value }

func (s SortByDecending) Len() int           { return len(s) }
func (s SortByDecending) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s SortByDecending) Less(i, j int) bool { return s[i].Salary.Value > s[j].Salary.Value }

func (p GroupByCurrency) GroupByCurrencyFunc() CurrencyGroups {
	currencyGroups := make(CurrencyGroups)
	for _, person := range p {
		currencyGroups[person.Salary.Currency] = append(currencyGroups[person.Salary.Currency], person)
	}
	return currencyGroups
}

func (f FilterBySalary) FilterBySalaryFunc(filter float64) []Person {
	var filtered []Person
	for _, person := range f {
		if person.Salary.Currency != "USD" {
			pair := fmt.Sprintf(`%s-USD`, person.Salary.Currency)
			//get exchange rate
			response := GetExchange(person.Salary.Currency)

			//convert salary to USD
			value := person.Salary.Value * response[pair]
			person.Salary.Value = value
			person.Salary.Currency = "USD"

			if value > filter {
				filtered = append(filtered, person)
			}
		} else if person.Salary.Value > filter {
			filtered = append(filtered, person)
		}

	}
	return filtered
}

func GetExchange(currency string) map[string]float64 {
	URL := fmt.Sprintf(`%s%s/exchange-rate`, os.Getenv("LOCAL_URL"), os.Getenv("PORT"))

	var pair = fmt.Sprintf(`%s-USD`, currency)
	data := map[string]string{"currency-pair": pair}
	var datajson, _ = json.Marshal(data)
	response, err := utils.HttpPOST(URL, "application/json", datajson)

	if err != nil {
		fmt.Println(err)
	}

	var exchangeRateResponse map[string]float64
	json.Unmarshal(response, &exchangeRateResponse)

	return exchangeRateResponse
}
