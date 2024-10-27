package arraysSlices

import "fmt"

// slice - do not have a pre-definend dimension of elements
// arrys - has a n number of dimensions
//var table []int
var table [10]int

// asign values to our array
var table1 [10]int = [10]int{10, 2, 0, 0, 10}

// with a matrix
var matrix [2][3]int

func ShowArray() {

	fmt.Println("Array Section: ")

	table[5] = 10
	table[9] = 10
	fmt.Println("This is the first Table")
	fmt.Println(table)

	fmt.Println("This is the table 2:")
	fmt.Println(table1)

	fmt.Println("Print the table1 using a for:")
	for i := 0; i < len(table1); i++ {

		fmt.Println(table1[i])
	}

	matrix[0][0] = 1
	matrix[1][1] = 1
	matrix[1][2] = 1
	fmt.Println(matrix)

}
