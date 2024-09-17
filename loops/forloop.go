package loops

import "fmt"

func PerformLoop() {

	// to perform a for loop for key - to break the foor loop - break
	//for {
	//	fmt.Println("Hello from a loop")
	//	break
	//}
	fmt.Println("Print values in a loop:")

	// we can perform
	i := 1
	for i <= 10 {

		fmt.Println(i)
		i++
	}

	fmt.Println("This is the j loop:")

	// optimize the previous
	for j := 0; j <= 100; j += 5 {

		fmt.Println(j)
	}

	// optimize the previous
	for j := 100; j >= 10; j -= 5 {

		fmt.Println(j)
	}

}
