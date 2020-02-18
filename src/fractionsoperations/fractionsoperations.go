package fractionsoperations

type FractionsOperations struct {
	result string
}

func (f FractionsOperations) validate(operation string) bool {
	return true;
}

func (f FractionsOperations) execute(operation string) string{
	return "execute"
}

func (f FractionsOperations) doOperation(operation string) string {
	return "doOperation"
}

func (f FractionsOperations) displayResult(operation string) string {
	return "displayResult"
}

func (f FractionsOperations) seachOperator(operation string) int {
	return 666
}

func (f FractionsOperations) newArray(operation string) string {
	return "newArray"
}