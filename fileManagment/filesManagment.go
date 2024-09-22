package fileManagment

import (
	"bufio"
	"fmt"
	"github.com/santos95/GoBasics/Exercises"
	"os"
)

// the route for the file to gererate
var fileName string = "./fileManagment/files/multiBoard.txt"

// function to write into a file
func WriteMultiChartIntoFIle() {

	// get the string to write
	strtwrt := Exercises.GetNumberMultiplicationChartText()

	// create a file - if exists is deleted and re-created
	file, err := os.Create(fileName)

	if err != nil {

		fmt.Println("Error during creation of the file " + err.Error())
		return
	}

	// write into the file
	fmt.Fprintln(file, strtwrt)

	// close the file - must be closed at the end
	file.Close()
}

func AppendMultiChartIntoFile() {

	// get the string to write
	strtwrt := Exercises.GetNumberMultiplicationChartText()

	// append into the file
	if !append(fileName, strtwrt) {

		fmt.Println("Error during append into file operation!")
	}

}

func append(filen string, text string) bool {

	// open the file - for this case append mode - with | we allow two send two flags
	// the | allow to pass more than one parameter
	// the two flags - the first one is write/read and the secoond oppend to append
	// the third parameter - permissions like linux permissions - wr to owner, read to group and others users
	// this allow to open the file into wr and append
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {

		fmt.Println("Error append text into the file ", err.Error())
		return false
	}

	// write into the file
	_, err = file.WriteString(text)

	if err != nil {

		fmt.Println("Error during write operation into file ", err.Error())
		return false
	}

	// close the file
	err = file.Close()

	if err != nil {

		fmt.Println("Error during close file operation ", err.Error())
		return false
	}

	return true
}

// to read the file and prompt into the command line
func ReadFile() {

	file, err := os.Open(fileName)

	if err != nil {

		fmt.Println("Error during read operation ", err.Error())
		return
	}

	scanner := bufio.NewScanner(file)

	// iterates over the file lines - when find a line return true else false
	for scanner.Scan() {

		record := scanner.Text()
		fmt.Println("> ", record)
	}

	err = file.Close()

	if err != nil {

		fmt.Println("Error during close file operation ", err.Error())
	}
}
