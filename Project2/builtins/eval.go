package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// prints error message for wrong number of arguments given, need at least one
var (
	ArgCounter = errors.New("invalid argument count")
)

func EvalComm(args ...string) error {
	switch len(args) {
	case 0: //need AT LEAST ONE ARGUMENT in order to be able to exectute the command, otherwise return error
		return fmt.Errorf("%w: at least one argument expected (command and its arguments)", ArgCounter)
	default:
		//Concatenates all given arguments into one string to create an executable command
		cmdString := strings.Join(args, " ")  //concatenates
		cmdParts := strings.Fields(cmdString) //creates command string

		// Executes the command based on the command string given
		cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("Could not execute the command: %w", err) //in case of error running the command, prints error message
		}
	}

	return nil
}

//example command used to test: ls -l to print
