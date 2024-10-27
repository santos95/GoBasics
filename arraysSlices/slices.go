package arraysSlices

import "fmt"

// for slices is the same as arrays but without predefined size
var tableSlice[]int = []int{1,2,3,4,5,6,7,8}

func ShowSlices() {

	fmt.Println("Slice table:")
	fmt.Println(tableSlice)
}