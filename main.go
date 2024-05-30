package main

import (
	"os"
)

func main() {
	var rotStr string
	
	if len(os.Args) != 3 {
		printLn("This program only takes one argument as input")
		printLn("example1: $go run . --to-rot14 \"input string\"")
		printLn("example2: $go run . --from-rot14 \"input string\"")
		return
	}

	arg := os.Args[1:3]

	if arg[0] == "--to-rot14" {
		rotStr = toRot14(arg[1])
	}

	

	printLn(rotStr)
}
