package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	//create a file
	content := "Lorem Ispam"
	file, err := os.Create("./myFile.txt")
	getError(err)

	//write in the aboe file
	length, err := io.WriteString(file, content)
	getError(err)

	fmt.Println("lenght is:", length)
	defer file.Close()
	readFile("./myFile.txt")

}

// define a method to read file
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	getError(err)
	fmt.Println("text data inside the file is:", string(databyte))
}

func getError(err error) {
	if err != nil {
		panic(err)
	}
}
