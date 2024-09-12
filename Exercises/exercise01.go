package Exercises

import (
	"fmt"
	"strconv"
)

func ExerciseStringConverter(stringValue string) (int, string) {

	var stringOutput string
	intOutput, _ := strconv.Atoi(stringValue)

	if intOutput > 100 {

		stringOutput = fmt.Sprintf("The value is bigger than 100!")

	} else {

		stringOutput = fmt.Sprintf("The value is less/equal than 100!")
	}

	return intOutput, stringOutput
}
