package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var num1 int
var num2 int
var text string
var err error

func InputNumbers() {

	// bufio new scanner - read from multiple sources, scanners, keyboards ..
	// we pass the os standar input - which is the keyboard
	// the standar output - console

	scanner := bufio.NewScanner(os.Stdin)

	// anythin we get with buffeo is a string - if we want to get a number we have to convert

	// check if the user promps a value throug the keyboard
	fmt.Println("Promp num1: ")
	if scanner.Scan() {

		// the text propertie - allows us to get the users input
		num1, err = strconv.Atoi(scanner.Text())

		if err != nil {
			panic("The input data is incorrect! " + err.Error())
		}

	}

	fmt.Println("Promp num2: ")
	if scanner.Scan() {

		num2, err = strconv.Atoi(scanner.Text())

		if err != nil {

			panic("The input data is incorrect! " + err.Error())
		}

	}

	fmt.Println("Input the text: ")
	if scanner.Scan() {

		text = scanner.Text()
	}

	fmt.Println(text, num1*num2)
}
