package Exercises

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// creates a function that promps for a number, if gets an error input promps to re-enter the number again.
func GetNumber() {

	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error

	for {

		fmt.Println("Enter a number: ")

		if scanner.Scan() {

			num, err = strconv.Atoi(scanner.Text())

			if err != nil {

				// if the input is incorrect, props for a correct input
				fmt.Println("Input incorrect, try it again!")
				continue

			} else {
				// if the input is correct continue with the print of the multiplication chart
				break
			}
		}
	}

	fmt.Println("The multiplication chart for the number", num, "is:")
	for i := 1; i <= 10; i++ {

		fmt.Println(num, " x ", i, " = ", num*i)
	}
}

func GetNumberMultiplicationChartText() string {

	scanner := bufio.NewScanner(os.Stdin)
	var num int
	var err error
	var multiChartText string

	for {

		fmt.Println("Enter a number: ")

		if scanner.Scan() {

			num, err = strconv.Atoi(scanner.Text())

			if err != nil {

				// if the input is incorrect, props for a correct input
				fmt.Println("Input incorrect, try it again!")
				continue

			} else {
				// if the input is correct continue with the print of the multiplication chart
				break
			}
		}
	}

	multiChartText += fmt.Sprintf("The multiplication chart for the number %d is:\n", num)
	for i := 1; i <= 10; i++ {

		multiChartText += fmt.Sprintf("%d x %d = %d \n", num, i, num*i)
	}

	return multiChartText
}
