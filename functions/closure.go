package functions

import "fmt"

// closure - functions that returns a functions
// multichar get a int value and return a function that returns a int
// Closures are functions that capture the values from their surrounding scope.
// allows to retain state between functions call without need to use a global variable
// the closures can remember the value of their enclose variables even when te functions has returned

// this functions return another function (the closure) that increments the sequence
func multiChar(value int) func() int {

	num := value
	sequence := 0 // captured by the closure - reteins his value between functions call

	// the actual closure - the functions that increment the value
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

//Closures are a powerful feature in Go that allow for state retention and encapsulation of logic. They are particularly useful in scenarios requiring callbacks or when you want to maintain some context across function calls.