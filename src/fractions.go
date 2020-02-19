package main

import (
	"fmt"
	"fractions/src/operations"
	_ "fractions/src/operations"
	"os"
	/*
	"strings"
	"strconv"
	"reflect"
	 */
)

func main() {
	operationArg := os.Args
	fractionsOperations := new(operations.FractionsOperations)

	if len(operationArg) > 1 && fractionsOperations.IsValid(operationArg[1]){
			fmt.Println("VALIDATE")
	} else {
		fmt.Println("Specify a valid operation to be evaluated")
	}







/*
	fmt.Println("MIERDA", operationArg);
	fmt.Println("---");
	fmt.Println("TYPE MIERDA", reflect.TypeOf(operationArg))
	fmt.Println("SIZE MIERDA", len(operationArg))

	operationArg = strings.Split(operationArg[1], " ")
	for i:=0; i< len(operationArg); i++ {
 		fmt.Println("VALUE " + strconv.Itoa(i) + " => ", operationArg[i])
	}
	fmt.Println("COMPLETE => ", operationArg)
*/
}