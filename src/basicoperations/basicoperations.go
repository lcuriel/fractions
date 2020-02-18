package basicoperations

type BasicOperations struct {
	numerator int
	denominator int
}

func (bo BasicOperations) multiplication(element1, element2 string) string {
	return "multiplication"
}

func (bo BasicOperations) division(element1, element2 string) string {
	return "division"
}

func (bo BasicOperations) addition(element1, element2 string) string {
	return "addition"
}

func (bo BasicOperations) subtraction(element1, element2 string) string {
	return "subtraction"
}

func (bo BasicOperations) reduce(element1, element2 string) string {
	return "reduce"
}

func (bo BasicOperations) integerToFraction(element1, element2 string) string {
	return "integerToFraction"
}
