package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	args := os.Args
	args = strings.Split(args[1], " ")
	for i:=0; i< len(args); i++ {
 		fmt.Println("VALUE " + strconv.Itoa(i) + " => ", args[i])
	}
	fmt.Println("COMPLETE => ", args)
}