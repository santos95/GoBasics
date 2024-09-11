package main

import (
	"fmt"
	"github.com/santos95/GoBasics/Variables"
)

func main() {

	fmt.Println("Hello, World!")
	fmt.Println("My name is Santos, I am learning Go ... !!!")

	// chapter 2 - variables and declare
	fmt.Println("Print variables values:")
	Variables.ShowIntegers()

	// chapter 3 - Continue variables and declare
	fmt.Println("Print different type of variables:")
	Variables.PrintVariables()

	// chapter 4 - functions
	state, text := Variables.ConvertToText(200)
	fmt.Println("State: ", state)
	fmt.Println("Converted: ", text)
}
