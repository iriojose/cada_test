package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"

	"github.com/iriojose/cada_test/types"
	"github.com/iriojose/cada_test/utils"
)

func GetPersonsHandler(w http.ResponseWriter, r *http.Request) error {
	var persons []types.Person

	FILEPATH := os.Getenv("FILEPATH")
	byteValue, err := utils.ReadFile(FILEPATH)
	if err != nil {
		return err
	}

	err = json.Unmarshal(byteValue, &persons)
	if err != nil {
		return err
	}

	if len(persons) == 0 {
		return fmt.Errorf("No persons found in file")
	}

	var ByAscendingSalary = make(types.SortByAscending, len(persons))
	var ByDecendingSalary = make(types.SortByDecending, len(persons))

	//copy persons va to ByAscendingSalary and ByDecendingSalary
	copy(ByAscendingSalary, persons)
	copy(ByDecendingSalary, persons)

	//make the sort
	sort.Sort(types.SortByAscending(ByAscendingSalary))
	sort.Sort(types.SortByDecending(ByDecendingSalary))

	//group persons by currency
	var groupByCurrency types.GroupByCurrency = persons
	var currencyGroups = groupByCurrency.GroupByCurrencyFunc()

	var ByFilter = types.FilterBySalary(persons).FilterBySalaryFunc(10000)

	//return diferent responses based on the exercise
	return utils.WriteJson(w, http.StatusOK, types.GetPersonsResponse{
		Persons:           persons,
		ByAscendingSalary: ByAscendingSalary,
		ByDecendingSalary: ByDecendingSalary,
		CurrencyGroups:    currencyGroups,
		ByFilter:          ByFilter,
	})
}
