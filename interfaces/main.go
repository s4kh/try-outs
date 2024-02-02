package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// s := square{sideLen: 4}
	// t := triangle{height: 5, base: 3}

	// printArea(s)
	// printArea(t)

	if len(os.Args) < 2 {
		log.Fatalln("please provide a filename. Ex. go run main.go test.txt")
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalln("error during open file: ", err)
	}

	io.Copy(os.Stdout, file)
	fmt.Printf("\nProgram executed successfully\n")
}
