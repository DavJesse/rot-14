package main

import (
	"os"
)

func main() {
	var rotStr string
	var inputStr string
	var flag string

	// Exit program when user inputs excess or insufficient arguments as input
	if len(os.Args) != 3 {
		printLn("This program only takes one argument as input")
		printLn("example1: $go run . --to-rot14 \"input string\"")
		printLn("example2: $go run . --from-rot14 \"input string\"")
		return
	}

	arg := os.Args[1:3]

	// Identify flag and inputstring
	if stringContains(arg[0], "rot14") {
		flag = arg[0]
		inputStr = arg[1]

		if flag == "--to-rot14" {
			rotStr = toRot14(inputStr)
		} else if flag == "--from-rot14" {
			rotStr = fromRot14(inputStr)
		} else {

			// Exit program when inompatible arguments are found
			printLn("Format Error!")
			printLn(flag + " is not a valid flag")

			return
		}

	} else if stringContains(arg[1], "rot14") {

		flag = arg[1]
		inputStr = arg[0]

		if flag == "--from-rot14" {
			rotStr = fromRot14(inputStr)
		} else if flag == "--to-rot14" {
			rotStr = toRot14(inputStr)
		} else {

			printLn("Format Error")
			printLn(flag + " is not a valid flag")

			return
		}

	} else {

		// Exit program when inompatible arguments are found
		printLn("Please include a '--to-rot14' or '--from-rot14' flag alongside your input string")
		printLn("example1. \"go run . --to-rot14 \"Your input string\"")
		printLn("example2. \"go run . --from-rot14 \"Your input string\"")

		return
	}

	// Print result of conversion to terminal
	printLn(rotStr)
}
