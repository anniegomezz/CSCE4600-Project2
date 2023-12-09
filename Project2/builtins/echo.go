package builtins

import (
	//manipulates errors
	"fmt"
	"strings"
)

// echo built-in command
// takes in strings and prints them to the terminal window
func EchoString(args ...string) error {
	//switch case in order to work with the input commands
	switch len(args) {
	case 0:
		//no arguments = prints empty line
		fmt.Println()
	case 1:
		//one argument = takes in the first argument and prints it
		fmt.Println(args[0])
	default:
		//multiple args = concatenate and print with space in between
		fmt.Println(strings.Join(args, " ")) //use join to join the arguments together, separated by a space
	}

	return nil //return if error
}
