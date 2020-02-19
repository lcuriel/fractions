package operations

import (
	"fmt"
	"regexp"
	"strings"
	_ "reflect"
)


type FractionsOperations struct {
	Result string
	Tracking []string
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
	elements := strings.Split(operation, " ")
	fmt.Println("ELEMENTS: ", elements)
	operators := [4]string{"*", "/", "+", "-"}
	for i:=0; i< len(operators); i++{
		fmt.Println("ELEMENT: ", operators[i])
		for {
			indexOperator := f.indexOf(elements, operators[i])
			element1, element2 := f.getElements(indexOperator)

			fmt.Println("INDEX: ", indexOperator)
			fmt.Println("==========")
			break
		}
	}
	/* */
	return "execute"
}

func (f FractionsOperations) DoOperation(operation string) string {
	return "doOperation"
}

func (f FractionsOperations) DisplayResult(operation string) string {
	return "displayResult"
}

func (f FractionsOperations) SeachOperator(operation string) int {
	return 666
}

func (f FractionsOperations) NewArray(operation string) string {
	return "newArray"
}

func (f FractionsOperations) indexOf(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
func (f FractionsOperations) getElements(index int) (string, string) {
	
	return "element01", "element02"
}




















