package Exercises

import (
	"strconv"
)

func ExerciseStringConverter(stringValue string) (int, string) {

	num, err := strconv.Atoi(stringValue)

	if err != nil {

		return 0, "There were a mistake during the convertion! - " + err.Error()
	}

	if num > 100 {

		return num, "The value is bigger than 100!"

	} else {

		return num, "The value is less/equal than 100!"
	}
}
