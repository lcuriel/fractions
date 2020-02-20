package operations

import (
	"fmt"
	"strconv"
	"strings"
	"reflect"
)

type BasicOperations struct {
	Numerator int
	Denominator int
}

func (bo BasicOperations) Multiplication(element1 string, element2 string) string {
	return "Multiplication"
}

func (bo BasicOperations) Division(element1 string, element2 string) string {
	return "Division"
}

func (bo BasicOperations) Addition(element1 string, element2 string) string {
	return "Addition"
}

func (bo BasicOperations) Subtraction(element1 string, element2 string) string {
	return "Subtraction"
}

func (bo BasicOperations) Reduce(element1 string, element2 string) string {
	return "Reduce"
}

func (bo BasicOperations) IntegerToFraction(element string) string {
	fmt.Println("**************************************************")
	fmt.Println("**************************************************")
	fmt.Println("**************************************************")
	fmt.Println("**************************************************")
	fmt.Println("**************************************************")
	fmt.Println("IntegerToFraction -> ORIGINAL: ", element)

	partsFraction := strings.Split(element, "_")
	if len(partsFraction) == 2 && len(partsFraction[1]) == 0 {
		partsFraction = []string{partsFraction[0]}
	}
	fmt.Println("IntegerToFraction -> partsFraction: ", partsFraction)
	fmt.Println("IntegerToFraction -> partsFraction -> LEN: ", len(partsFraction))

	// element == 3/4
	if strings.Contains(partsFraction[0], "/") {
		fmt.Println("IntegerToFraction -> FRACTION: ", partsFraction[0])
		return element
	}
	// element == 5_
	if len(partsFraction) == 1 {
		fmt.Println("IntegerToFraction -> INTEGER: ", partsFraction[0]+"/1")
		return  partsFraction[0]+"/1"
	}
	// element == 5_3/4
	if len(partsFraction) == 200  {
		xtem, _ := strconv.Atoi(partsFraction[0])
		fmt.Println("=============================================")
		fmt.Println("DATA: ", xtem)
		fmt.Println("TYPE: ", reflect.TypeOf(xtem))
		fmt.Println("=============================================")
		fmt.Println("")
		fmt.Println("")



		//return strconv.Itoa( * denominator) + "/2"
	}

	// element == 5_
	// element == 5_3/4


	/*
	// [5_], [5_ 3/4]





	if !strings.Contains(partsFraction[0], "/") || len(partsFraction) == 2 {
		integer, _ := strconv.Atoi(partsFraction[0])
	}
	// [3/4], [5_ 3/4]
	if strings.Contains(partsFraction[0], "/") || len(partsFraction) == 2 {
		element = partsFraction[0] //[3/4]
		if len(partsFraction) == 2 {
			element = partsFraction[1] //[5_ 3/4]
		}
		partsFraction = strings.Split(element, "/")
		numerator, _ := strconv.Atoi(partsFraction[0])
		denominator, _ := strconv.Atoi(partsFraction[1])
	}
*/


	return "integerToFraction"
}
