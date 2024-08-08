package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in go lang")

	content := "this needs to go in files - majidbeg"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	checkNilErr(err)

	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(fileName string) {
	databyte, err := os.ReadFile(fileName)
	checkNilErr(err)
	fmt.Println("text data inside the file is\n : ", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
