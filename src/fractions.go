package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
	"reflect"
	fo "fractions/src/fractionsoperations"
)

func main() {
	operationArg := os.Args
	var fos fractionsoperations.FractionsOperations
	fmt.Println("VALIDATE",   )










	fmt.Println("MIERDA", args);
	fmt.Println("---");
	fmt.Println("TYPE MIERDA", reflect.TypeOf(args))
	fmt.Println("SIZE MIERDA", len(args))

	args = strings.Split(args[1], " ")
	for i:=0; i< len(args); i++ {
 		fmt.Println("VALUE " + strconv.Itoa(i) + " => ", args[i])
	}
	fmt.Println("COMPLETE => ", args)

}