package arraysSlices

import "fmt"

// slice - do not have a pre-definend dimension of elements
// arryyas - has a n number of dimensions
//var table []int
var table [10]int

// asign values to our array
var table1 [10]int = [10]int{10, 2, 0, 0, 10}

func ShowArray() {

	fmt.Println("Array Section: ")

	table[5] = 10
	table[9] = 10
	fmt.Println("This is the first Table")
	fmt.Println(table)

	fmt.Println("This is the table 2:")
	fmt.Println(table1)
}
