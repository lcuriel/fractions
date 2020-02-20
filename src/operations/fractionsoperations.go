package operations

import (
	"fmt"
	"regexp"
	"strings"
	_ "reflect"
)


type FractionsOperations struct {
	result string
	tracking []string
	elements []string
}

func (f *FractionsOperations) IsValid(operation string) bool {
	band := false
	operation = f.CleanString(operation)
	if len(operation) > 1 {
		expression := `^(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))(\s(\*|\/|\+|\-)\s(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*)))*$`
		rsEvaluator, _ := regexp.Compile(expression)
		band = rsEvaluator.MatchString(operation)
		fmt.Println("PATRON: ", band)
	}
	return band
}

func (f FractionsOperations) CleanString(operation string) string{
	operation = strings.Trim(operation, " ")
	for{
		spaces := strings.Contains(operation, "  ")
		if !spaces {
			break
		}
		operation = strings.Replace(operation, "  ", " ", -1)
	}
	return operation
}

func (f FractionsOperations) Execute(operation string) string{
	result := ""
	f.elements = strings.Split(operation, " ")

	operators := [4]string{"*", "/", "+", "-"}
	for i:=0; i< len(operators); i++{
		/*
		fmt.Println("OPERATOR: ", operators[i])
		fmt.Println("=========================")
		 */
		for {
			indexOperator := f.indexOf(f.elements, operators[i])
			if indexOperator >= 0 {
				element1, element2 := f.getElements(indexOperator)
				result = f.doOperation(operators[i], element1, element2)
				f.newArray(f.elements, result)
			}
			break
		}
		i = len(operators)
	}
	return "execute"
}

func (f FractionsOperations) getElements(indexOperator int) (string, string) {
	fmt.Println("ELEMENTS: ", f.elements)
	fmt.Println("indexOperator: ", indexOperator)
	e1 := f.elements[indexOperator-1]
	e2 := f.elements[indexOperator+1]
	return e1, e2
}

func (f FractionsOperations) doOperation(operator string, element1 string, element2 string) string {
	bo := new(BasicOperations)
	var result string
	element1 = bo.IntegerToFraction(element1)
	element2 = bo.IntegerToFraction(element2)
	switch operator{
		case "*":
			result = bo.Multiplication(element1, element2)
		case "/":
			result = bo.Division(element1, element2)
		case "+":
			result = bo.Addition(element1, element2)
		case "-":
			result = bo.Subtraction(element1, element2)
	}
	return result
}

func (f FractionsOperations) indexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
/*
 ===================================================================================================
 ===================================================================================================
 ===================================================================================================
 ===================================================================================================
 */
func (f FractionsOperations) newArray(operation []string, result string) string {
	return "newArray"
}

func (f FractionsOperations) DisplayResult(operation string) string {
	return "displayResult"
}





















