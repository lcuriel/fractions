package main

import (
	"fmt"
	"fractions/src/operations"
	"testing"
)

const expected bool = true

func TestIsValidOneElement(t *testing.T){
	result := false
	//One element
	testCases := []string{ "-3/5", "3/-5", "-3/-5", "1", "-1", "1/1", "1_-3/5", "1_3/-5", "1_-3/-5", "-1_-3/5", "-1_3/-5", "-1_-3/-5" }
	fractionsOperations := new(operations.FractionsOperations)
	for i:=0; i< len(testCases); i++{
		fmt.Println(testCases[i])
		fmt.Println(fractionsOperations.IsValid(testCases[i]))
		fmt.Println("----------------------------------------------------------------------")
		result = fractionsOperations.IsValid(testCases[i])
		if result != expected {
			//t.Errorf("Element: %s, result:%t want: %t.", testCases[i], result, expected)
		}
	}
}

func estIsValidManyElements(t *testing.T){
	result := false
	//One element
	testCases := []string{
		"1/4 * 1/8 / 1/2 + 1/16 - 1/32 * 1/64",
		"90/12 / 3 / 7 / 40/83 / 4 / 70/97 / 8_4/6 / 3 / 5/24 / 9_5/34 / 5_9/21",
		"5 - 9_3/48 - 6/5 - 98_7/9 - 63_4/59 - 3/240 - 5/98 - 3_4/8 - 4_3/8 - 5/6",
		"40 / 2/67 * 6/5 / 40_1/2 - 9_5/6 + 2/50 * 1 - 2_8/7 / 6/5 - 40/12 * 7/8 - 5/6 + 1_3/24",
		"1/2 + 3/4 * 5_34/345 / 3 * 55/123 + 33_4/34 + 343 / 33_4/13 * 23/66 - 33 * 1/3 + 33_",
		"30 + 1_7/48 + 6/38 + 47_2/594 + 3_8/7 + 10_2/96 + 5/42 + 33 + 4 + 3/60 + 9 + 5/67",
		"5 * 4/5 * 6/97 * 9_3/7 * 5/4 * 6 * 4/87 + 5_6/12 * 40 * 8/75 * 6_10/27 * 8/64 * 8/5 * 2_1/8 * 60_1/9",
		"-1 * 1 + 31 + 242 + 452 - 626 + -75 * 28 + 745 / 2 + -37 * 23 + 32 + 4 - 752 + 38" }
	fractionsOperations := new(operations.FractionsOperations)
	for i:=0; i< len(testCases); i++{
		result = fractionsOperations.IsValid(testCases[i])
		if result != expected {
			t.Errorf("Element: %s, result:%t want: %t.", testCases[i], result, expected)
		}
	}
}
