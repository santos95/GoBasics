package main

import (
	"fmt"
	"github.com/santos95/GoBasics/Exercises"
	"github.com/santos95/GoBasics/Variables"
	"github.com/santos95/GoBasics/fileManagment"
	"github.com/santos95/GoBasics/functions"
	"github.com/santos95/GoBasics/loops"
	"runtime"
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

	// chapter 5 - control structures
	// the runtime package allow us to get info of our enviroment - so, arquitecture, enviroment variables..

	// if structure - allow us to assing and inmediatly evaluate as follows:

	if os := runtime.GOOS; os == "Linux" {

		fmt.Println("This is a linux enviroment")
	} else {

		fmt.Println("This is a ", os, " enviroment.")
	}

	// we can only evaluate if os == "Linux" {}

	// switch structure
	fmt.Println("Switch structure:")
	switch os := runtime.GOOS; os {

	case "Linux":
		fmt.Println("This is a linux environment!")
	case "Darwin":
		fmt.Println("This is a Darwin envirnment")
	default:
		fmt.Printf("This is a %s  environment!\n", os)
	}

	// exercise 1
	num, stringOut := Exercises.ExerciseStringConverter("20")
	fmt.Println("Result1: ", num, ", ", stringOut)

	// chapter6 input/output from console
	//keyboard.InputNumbers()

	// chapter7 loops
	loops.PerformLoop()

	// exercise 2
	Exercises.GetNumber()

	// chapter 8 - file managment
	fmt.Println("This is the file managment output:")
	//fmt.Println(Exercises.GetNumberMultiplicationChartText())
	//fileManagment.WriteMultiChartIntoFIle()
	// comment to avoid adding more data to the file for now
	//fileManagment.AppendMultiChartIntoFile()

	// read and print the file into the command line
	fmt.Println("This is file managment read file:")
	fileManagment.ReadFile()

	// chapter 9 - anonymous functions and closures
	functions.MathOperation()
	fmt.Println("Closure use:")
	functions.CallClosure()

	// chapter 10 - recursion
	functions.PowerOf()
}
