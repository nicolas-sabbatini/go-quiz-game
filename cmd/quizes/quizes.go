package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	// os.Args is a slice of strings
	// os.Args[0] is the name of the command
	// os.Args[1:] are the arguments
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1:])
}
