package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	ErrInvalidArgCounter = errors.New("invalid argument count")
)

func PwdComm(args ...string) error {
	switch len(args) {
	case 0:
		//Gets the current working directory
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("Error. Unable to fetch current working directory: %w", err)
		}
		fmt.Println(cwd) //If no error, prints the current working directory
	default:
		return fmt.Errorf("%w: expected zero arguments", ErrInvalidArgCounter)
	}

	return nil
}
