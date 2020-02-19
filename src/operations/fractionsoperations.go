package operations

import (
	"fmt"
	"regexp"
	_ "bytes"
	"strings"
)


type FractionsOperations struct {
	Result string
}

func (f FractionsOperations) IsValid(operation string) bool {
	band := false
	operation = f.cleanString(operation)
	if len(operation) > 1 {
		expression := `^(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*))(\s(\*|\/|\+|\-)\s(([1-9]\d*\_)|([1-9]\d*\/[1-9]\d*)|([1-9]\d*\_[1-9]\d*\/[1-9]\d*)))*$`
		rsevaluator, _ := regexp.Compile(expression)
		band = rsevaluator.MatchString(operation)
		fmt.Println("PATRON: ", band)
	}
	return band
}

func (f FractionsOperations) cleanString(operation string) string{
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