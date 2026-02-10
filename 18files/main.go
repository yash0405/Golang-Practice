package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("File Handling")

	content := "This needs to go in a file"

	file, err := os.Create("./mytestfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("Length is: ", length)
	defer file.Close()

	readFile("./mytestfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Text data inside the file is: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
