package main

import (
	"fmt"
	"fractions/src/operations"
	"os"
)

func main() {
	operationArg := os.Args
	fractionsOperations := new(operations.FractionsOperations)

	if len(operationArg) > 1 {
		operation := fractionsOperations.CleanString(operationArg[1])
		operationArg = nil
		if  fractionsOperations.IsValid(operation){
			result := fractionsOperations.Execute(operation)
			println(result)
		} else {
			fmt.Println("Specify a valid operation to be evaluated")
		}
	} else {
		fmt.Println("Enter an operation")
	}
}