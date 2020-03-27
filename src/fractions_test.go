package main

import (
	"fractions/src/operations"
	"testing"
)

const expected bool = true

func TestIsValidOneElement(t *testing.T){
	result := false
	//One element
	tables := []struct{
		element string
		expected bool
	}{
		{"-3/5", true},
		{"3/-5", true},
		{"-3/-5", true},
		{"1", true},
		{"-1", true},
		{"1/1", true},
		{"1_-3/5", true},
		{"1_3/-5", true},
		{"1_-3/-5", true},
		{"-1_-3/5", true},
		{"-1_3/-5", true},
		{"-1_-3/-5", true},
	}
	fractionsOperations := new(operations.FractionsOperations)
	for _, table := range tables {
		result = fractionsOperations.IsValid(table.element)
		if result != table.expected {
			t.Errorf("Element: %s, result:%t want: %t.", table.element, result, expected)
		}
	}
}

func TestIsValidManyElements(t *testing.T){
	result := false
	tables := []struct{
		element string
		expected bool
	}{
		{ "1/4 * 1/8 / 1/2 + 1/16 - 1/32 * 1/64", true },
		{ "90/12                                           /             3     /      7      /           40/83             / 4             /          70/97                                              /                8_4/6 / 3             / 5/24                                     /  9_5/34    /       5_9/21", true },
		{ "5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6", true },
		{ "40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1 - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24", true },
		{ "1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33", true },
		{ "30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67", true },
		{ "5 * 4/5 * 6/97 * 9_3/7 * 5/4 * 6 * 4/87 + 5_6/12 * 40 * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9", true },
		{ "-1    *    1    +    31    +    242    +    452    - 626    +    -75    *    28    +    745    /    2    +    -37    *    23    +    32    +    4    -    752    +    38", true },
	}
	fractionsOperations := new(operations.FractionsOperations)
	for _, table := range tables {
		result = fractionsOperations.IsValid(fractionsOperations.CleanString(table.element))
		if result != table.expected {
			t.Errorf("Element: %s, result:%t want: %t.", table.element, result, expected)
		}
	}
}

func TestIsValidWrongSyntax(t *testing.T){
	result := false
	tables := []struct{
		element string
		expected bool
	}{
		{ "1/4 * 1/8 / 1/2 + 1/16 - 1/32 * 1/64", true },
		{ "90/12 / 3 / 7 / 40/83 / 4 / 70/97 / 8_4/6 / 3 / 5/24 / 9_5/34 / 5_9/21", true },
		{ "5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6", true },
		{ "40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1 - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24", true },
		{ "1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33", true },
		{ "30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67", true },
		{ "5 * 4/5 * 6/97 * 9_3/7 * 5/4 * 6 * 4/87 + 5_6/12 * 40 * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9", true },
		{ "-1 * 1 + 31 + 242 + 452 - 626 + -75 * 28 + 745 / 2 + -37 * 23 + 32 + 4 - 752 + 38", true },
	}
	fractionsOperations := new(operations.FractionsOperations)
	for _, table := range tables {
		result = fractionsOperations.IsValid(table.element)
		if result != table.expected {
			t.Errorf("Element: %s, result:%t want: %t.", table.element, result, expected)
		}
	}
}

func TestExecute(t *testing.T){
	result := ""
	tables := []struct{
		element string
		expected string
	}{
		{ "1/4 * 1/8 / 1/2", "1/16" },
		{ "90/12 / 3 / 7 / 40/83 / 4 / 70/97 / 8_4/6 / 3 / 5/24 / 9_5/34 / 5_9/21", "410601/430175200" },
		{ "3 + 3/4 - 4/5 + 7/9 - 4/5 + 3/4", "28/45" },
		{ "40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1 - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24", "10_69617/340200" },
		{ "1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33", "22_151883/8097100" },
		{ "30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67", "138_37817327/52931340" },
		{ "5 * 4 / 5 * 6 / 7 * 3", "2/63" },
		{ "20 - 1 + 31 + 1/2 + 1/2 - 1/6 + -75 - 28 + 1/5 - 2 + -37 - 23 + 32 + 4 - 1/2 + 3", "6_2/15" },
	}
	fractionsOperations := new(operations.FractionsOperations)
	for _, table := range tables {
		result = fractionsOperations.Execute(table.element)
		if result != table.expected {
			t.Errorf("Element: %s, result:%s want: %s.", table.element, result, table.expected)
		}
	}

}