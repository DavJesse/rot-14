package main

import (
	"os"
)

func main() {
	var rotStr string
	var inputStr string
	var flag string

	if len(os.Args) != 3 {
		printLn("This program only takes one argument as input")
		printLn("example1: $go run . --to-rot14 \"input string\"")
		printLn("example2: $go run . --from-rot14 \"input string\"")
		return
	}

	arg := os.Args[1:3]

	if stringContains(arg[0], "-rot14") {
		flag = arg[0]
		inputStr = arg[1]
		if stringContains(flag, "to") {
			rotStr = toRot14(inputStr)
		}
	} else if stringContains(arg[1], "-rot14") {
		flag = arg[1]
		inputStr = arg[0]
		if stringContains(flag, "from") {
			rotStr = fromRot14(inputStr)
		}

	} else {
		printLn("Please include a '--to-rot14' or '--from-rot14' flag alongside your input string")
		printLn("Ex. \"go run . --to-rot14 \"Your input string\"")
		printLn("Ex. \"go run . --from-rot14 \"Your input string\"")
		return
	}

	printLn(rotStr)
}
