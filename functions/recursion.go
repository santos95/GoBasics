package functions

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func PowerOf() {

	getNumber := func() int {

		scanner := bufio.NewScanner(os.Stdin)

		for {

			fmt.Println("Input a number: ")

			if scanner.Scan() {

				n, err := strconv.Atoi(scanner.Text())

				if err != nil {

					fmt.Println("Input incorrect ", err.Error())
					continue
				}

				return n
			}
		}
	}

	fmt.Println("Input the base number:")
	num1 := getNumber()
	fmt.Println("Input the times number:")
	num2 := getNumber()

	fmt.Println("The result is:")
	fmt.Println(getPower(num1, num1, num2))

}

func getPower(num1 int, num2 int, num3 int) int {

	if num3 == 1 {

		return num1
	}

	num3 -= 1
	return getPower(num1*num2, num2, num3)
}
