package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// welcome to your channel go guru

// Topic :- file system

// golang offers inbuilt lib to create file along with that read and write files.
// LIB LIKE  io/ioutil, os, bufio

// methods :-
// os.Create()
// ioutil.ReadFile()
// ioutil.WriteFile()
// bufio.NewReader(os.Stdin)

func CreateFile() {

	fmt.Println("Writing to a file in Go lang")

	file, err := os.Create("abc.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		// os.Exit
	}

	defer file.Close()

	len, err := file.WriteString("asdfsdsdfsdffsfsdfsWelcome to your channel go guru. :)")

	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}

func ReadFile() {

	fmt.Printf("\n\nReading a file in Go lang\n")
	fileName := "fghdfg.txt"
	data, err := ioutil.ReadFile("fgdhdf.txt")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)

}

// main function
func main() {

	CreateFile()
	// ReadFile()
}
