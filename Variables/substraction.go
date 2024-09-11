package Variables

import (
	"fmt"
	"strconv"
	"time"
)

var Name string
var State bool
var Salary float32
var Actualdate = time.Time{}

func PrintVariables() {

	Name = "Peter Parker"
	State = true
	Salary = 20.25
	Actualdate = time.Now()

	fmt.Println(Name)
	fmt.Println(State)
	fmt.Println(Salary)
	fmt.Println(Actualdate)
}

func ConvertToText(number int) (bool, string) {
	texto := strconv.Itoa(number)
	return true, texto
}
