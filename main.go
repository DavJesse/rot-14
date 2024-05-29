package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printLn("This program only takes one argument as input")
		printLn("In the case of multiple words, encapsulate in double or single quotes")
		return
	}

	arg := os.Args[1]

	rotStr := rot14(arg)

	printLn(rotStr)
}
