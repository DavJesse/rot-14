package main

import (
	"os"
)

func main() {
	if len(os.Args) != 2 {
		printLn("This program only takes one argument as input")
		return
	}

	arg := os.Args[1]

	rotStr := rot14(arg)

	printLn(rotStr)
}
