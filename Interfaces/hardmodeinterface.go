package main

// This program reads the content of a text file named "example.txt" located in the same directory
// and prints its content to the console. The program assumes the file exists and is readable.

import (
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
