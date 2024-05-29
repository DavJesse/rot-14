package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printLn("This program only takes one argument as input")
		return
	}

	arg := os.Args[1]

	fmt.Println(arg)
}
