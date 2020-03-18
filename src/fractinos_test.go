package main

import (
	"testing"
	"fractions/src/operations"
)



func TestIsvalid(t *testing.T){
	test := false
	expected := false
	testCases := { "" }
	fractionsOperations := new(operations.FractionsOperations


	if test != expected {
		t.Errorf("The test was incorrect, got: %t, want: %t.", test, expected)
	}
}