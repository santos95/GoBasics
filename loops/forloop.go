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

		// breaks the loop when j is equal to - abort the loop execution
		if j == 20 {
			break
		}
		fmt.Println(j)
	}

	// optimize the previous
	for j := 100; j >= 10; j -= 5 {

		// with the continue the condition is re-evaluted and avoid the execution of the subsequence instructions
		if j == 20 {
			continue
		}
		// when j == 20 this line is not executed
		fmt.Println(j)
	}

}
