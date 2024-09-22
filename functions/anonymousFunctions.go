package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// anonymous functions -
//functions without name that can be assigned
//and passed as parameter
func MathOperation() {

	var num1, num2 int
	num3 := 10

	getANumber := func() int {

		scanner := bufio.NewScanner(os.Stdin)

		for {

			fmt.Println("Number: ")
			if scanner.Scan() {

				num, err := strconv.Atoi(scanner.Text())

				if err != nil {
					fmt.Println("Input incorrect ", err.Error())
					continue
				}

				return num
			}
		}
	}

	// input number 1
	fmt.Println("Number 1")
	num1 = inputInteger()

	fmt.Println("Number 2")
	num2 = getANumber()

	// assing the function to the variable - function without name
	operation := func(n1 int, n2 int) int {

		return n1 + n2
	}

	// here we are sending the functions as parameter
	fmt.Println("The operation result is: ", operation(num1, num2))

	// we can modify the behavior of the function for anonymous functions
	// we can't change their structure - only the behavior
	// but this is for the var operation as a functions
	operation = func(n1 int, n2 int) int {

		return (n1 + n2) * num3
	}

	fmt.Println("The operation result is: ", operation(num1, num2))
}

func inputInteger() int {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Number: ")
		if scanner.Scan() {

			num, err := strconv.Atoi(scanner.Text())

			if err != nil {
				fmt.Println("Input incorrect ", err.Error())
				continue
			}

			return num
		}
	}
}
