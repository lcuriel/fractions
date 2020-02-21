package operations

import (
	"fmt"
	_ "fmt"
	"strconv"
	"strings"
)

type FractionDestructured struct {
	numerator int
	denominator int
}

type BasicOperations struct {
	element1 int
	element2 int
}

func (bo BasicOperations) IntegerToFraction(element string) string {
	partsFraction := strings.Split(element, "_")
	if len(partsFraction) == 2 && len(partsFraction[1]) == 0 {
		partsFraction = []string{partsFraction[0]}
	}
	if strings.Contains(partsFraction[0], "/") { 	// element ==> FRACTION (3/4, 1/2, etc)
		return element
	}
	if len(partsFraction) == 1 {	// element ==> INTEGER (5_, 666_, 999_)
		return  partsFraction[0]+"/1"
	}
	integer, _ := strconv.Atoi(partsFraction[0])	// element ==> INTEGER-FRACTION (5_3/4, 1_2/3, 1243_234234/233467)
	partsFraction = strings.Split(partsFraction[1], "/")
	numerator, _ := strconv.Atoi(partsFraction[0])
	denominator, _ := strconv.Atoi(partsFraction[1])
	numerator = (integer * denominator) + numerator
	newFraction := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
	return newFraction
}

func (bo BasicOperations) Multiplication(element1 string, element2 string) string {
	e1 := bo.fractionDestructuring(element1)
	e2 := bo.fractionDestructuring(element2)
	numerator := e1.numerator * e2.numerator
	denominator := e1.denominator * e2.denominator
	result := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
	return result
}

func (bo BasicOperations) Division(element1 string, element2 string) string {
	e1 := bo.fractionDestructuring(element1)
	e2 := bo.fractionDestructuring(element2)
	numerator := e1.numerator * e2.denominator
	denominator := e1.denominator * e2.numerator
	result := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
	return result
}

func (bo BasicOperations) Addition(element1 string, element2 string) string {
	e1 := bo.fractionDestructuring(element1)
	e2 := bo.fractionDestructuring(element2)
	numerator := (e1.numerator * e2.denominator) + (e1.denominator * e2.numerator)
	denominator := e1.denominator * e2.denominator
	result := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
	return result
}

func (bo BasicOperations) Subtraction(element1 string, element2 string) string {
	e1 := bo.fractionDestructuring(element1)
	e2 := bo.fractionDestructuring(element2)
	numerator := (e1.numerator * e2.denominator) - (e1.denominator * e2.numerator)
	denominator := e1.denominator * e2.denominator
	result := strconv.Itoa(numerator) + "/" + strconv.Itoa(denominator)
	return result
}

func (bo BasicOperations) fractionDestructuring(element string) FractionDestructured {
	partsFraction := strings.Split(element, "/")
	numerator, _ := strconv.Atoi(partsFraction[0])
	denominator, _ := strconv.Atoi(partsFraction[1])
	return FractionDestructured{ numerator: numerator, denominator: denominator }
}



















func (bo BasicOperations) Reduce(element string) string {
	fmt.Println("ORIGINAL: ", element)
	integer := 0
	result := ""
	e := bo.fractionDestructuring(element)
	if (e.numerator > e.denominator) || (e.numerator == e.denominator) {
		integer = int(e.numerator / e.denominator)
		e.numerator = e.numerator - (integer * e.denominator)
		result = strconv.Itoa(integer) + "_"
		if e.numerator == 0 {
			return result
		}
	}
	divisor := bo.findDivisor(e)
	e.numerator = e.numerator / divisor
	e.denominator = e.denominator / divisor
	result = result + strconv.Itoa(e.numerator) + "/" + strconv.Itoa(e.denominator)
	return result
}

func (bo BasicOperations) findDivisor(element FractionDestructured) int {
	small := element.numerator
	if element.denominator < element.numerator{
		small = element.denominator
	}
	result := 1
	for i:=small; i>0; i--{
		if (element.numerator%i == 0) && (element.denominator%i == 0) {
			result = i
			i = 0
			break
		}
	}
	return result
}

















