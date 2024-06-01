package main

import "os"

// Prints string parsed to commandline
func printLn(str string) {
	os.Stdout.WriteString(str + "\n")
}
