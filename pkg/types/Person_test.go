package types

import (
	"reflect"
	"sort"
	"testing"
)

func TestSortByAscending(t *testing.T) {
	people := SortByAscending{
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "USD"}},
	}

	sort.Sort(SortByAscending(people))

	expected := SortByAscending{
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "USD"}},
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
	}

	if !reflect.DeepEqual(people, expected) {
		t.Errorf("unexpected result, got: %v, want: %v.", people, expected)
	}
}

func TestSortBySortByDecending(t *testing.T) {
	people := SortByDecending{
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "USD"}},
	}

	sort.Sort(SortByDecending(people))

	expected := SortByDecending{
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "USD"}},
	}

	if !reflect.DeepEqual(people, expected) {
		t.Errorf("unexpected result, got: %v, want: %v.", people, expected)
	}
}

func TestGroupByCurrency(t *testing.T) {
	people := GroupByCurrency{
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "VES"}},
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "EUR"}},
	}

	groupByCurrency := people.GroupByCurrencyFunc()

	expected := CurrencyGroups{
		"USD": {
			{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		},
		"VES": {
			{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "VES"}},
		},
		"EUR": {
			{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 45000, Currency: "EUR"}},
		},
	}

	if !reflect.DeepEqual(groupByCurrency, expected) {
		t.Errorf("unexpected result, got: %v, want: %v.", people, expected)
	}
}

func TestFilterBySalary(t *testing.T) {
	people := FilterBySalary{
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
		{Id: 3, PersonName: "Charlie", Salary: Salary{Value: 10000, Currency: "EUR"}},
	}

	filtered := people.FilterBySalaryFunc(48000)

	expected := []Person{
		{Id: 1, PersonName: "Alice", Salary: Salary{Value: 50000, Currency: "USD"}},
		{Id: 2, PersonName: "Bob", Salary: Salary{Value: 60000, Currency: "USD"}},
	}

	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("unexpected result, got: %v, want: %v.", people, expected)
	}
}
