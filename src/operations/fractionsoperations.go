package operations

import (
	"fmt"
	"regexp"
	"strings"
)


type FractionsOperations struct {
	result string
	tracking []string
	elements []string
}

func (f *FractionsOperations) IsValid(operation string) bool {
	band := false
	operation = f.CleanString(operation)
	if len(operation) >= 1 {
		expression := `^((\-?[1-9]\d*)|(\-?[1-9]\d*\/\-?[1-9]\d*)|(\-?[1-9]\d*\_\-?[1-9]\d*\/\-?[1-9]\d*))(\s(\*|\/|\+|\-)\s((\-?[1-9]\d*)|(\-?[1-9]\d*\/\-?[1-9]\d*)|(\-?[1-9]\d*\_\-?[1-9]\d*\/\-? [1-9]\d*)))*$`
		rsEvaluator, _ := regexp.Compile(expression)
		band = rsEvaluator.MatchString(operation)
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

func (f *FractionsOperations) Execute(operation string) string{
	result := ""
	f.elements = strings.Split(operation, " ")
	bo := new(BasicOperations)
	if len(f.elements)==1 {
		result = bo.IntegerToFraction(f.elements[0])
		result = bo.Reduce(result)
	} else {
		operators := [4]string{"*", "/", "+", "-"}
		for i:=0; i< len(operators); i++{
			for {
				f.newEntryLog()
				indexOperator := f.indexOf(f.elements, operators[i])
				if indexOperator >= 0 {
					element1, element2 := f.getElements(indexOperator)
					result = f.doOperation(operators[i], element1, element2)
					f.newArray(indexOperator, bo.Reduce(result))
				} else {
					break
				}
			}
		}
		result = f.elements[0]
	}
	return result
}

func (f FractionsOperations) getElements(indexOperator int) (string, string) {
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

func (f *FractionsOperations) newEntryLog(){
	newString := strings.Join(f.elements, " ")
	f.tracking = append(f.tracking, newString)
}

func (f *FractionsOperations) GetLog() []string {
	return f.tracking
}

func (f *FractionsOperations) PrintLog(log bool) {
	if log {
		for i:=0; i < len(f.tracking); i++ {
			fmt.Println(i, " => ", f.tracking[i])
		}
	}
}

func (f *FractionsOperations) newArray(indexOperator int, newElement string) {
	indexOperatorM1 := indexOperator - 1
	indexOperatorP1 := indexOperator + 1
	f.elements[indexOperatorM1] = newElement
	f.elements = append(f.elements[:indexOperator], f.elements[indexOperatorP1:]...)
	f.elements = append(f.elements[:indexOperator], f.elements[indexOperatorP1:]...)
}
