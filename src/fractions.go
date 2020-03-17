package main

import (
	"fmt"
	"fractions/src/operations"
	"os"
)

func main() {
	operationArg := os.Args
	log := false
	fractionsOperations := new(operations.FractionsOperations)

	if len(operationArg) > 1 {
		operation := fractionsOperations.CleanString(operationArg[1])
		if len(operationArg) >= 3 && operationArg[2] == "logs" {
			log = true
		}
		operationArg = nil
		if  fractionsOperations.IsValid(operation){
			result := fractionsOperations.Execute(operation)
			println(result)
			fractionsOperations.PrintLog(log)
		} else {
			fmt.Println("Specify a valid operation to be evaluated")
		}
	} else {
		fmt.Println("Enter an operation")
	}
}