package functions

import "fmt"

// closure - functions that returns a functions
// multichar get a int value and return a function that returns a int
func multiChar(value int) func() int {

	num := value
	sequence := 0

	return func() int {
		sequence++
		return num * sequence
	}
}

func CallClosure() {

	multiChartFrom := 2
	myChar := multiChar(multiChartFrom)

	for i := 0; i <= 10; i++ {

		fmt.Println(myChar())
	}
}
