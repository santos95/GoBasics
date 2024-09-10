package Variables

import "fmt"

// to make a method public - start with a capatalize letter
func ShowIntegers() {
	//	declare variables - two ways - declare and assing
	var commonint int
	varint32 := int32(10) // returns a 32 bit integer - the var takes the type of the assing
	varint64 := int64(100)

	// the variables are initialized with the minimum value of the data type.
	// So, they are not initialized with null
	// commonint will have a value of 0, the minimal value.

	fmt.Println("commonint = ", commonint)
	fmt.Println("varint32 = ", varint32)
	fmt.Println("varint64 = ", varint64)

}
