package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	args := os.Args

	fmt.Println("MIERDA", args);
	fmt.Println("---");
	fmt.Println("---");
	fmt.Println("---");


	args = strings.Split(args[1], " ")
	for i:=0; i< len(args); i++ {
 		fmt.Println("VALUE " + strconv.Itoa(i) + " => ", args[i])
	}
	fmt.Println("COMPLETE => ", args)
}