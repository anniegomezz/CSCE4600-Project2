package builtins

import (
	"errors"
	"fmt"
	"syscall"
	"time"
)

var (
	NoArgsNeeded = errors.New("Incorrect usage. No arguments are needed for this command.")
)

// Times command that prints out the user and system times used by the shell and its children
func TimesComm(args ...string) error {
	switch len(args) {
	case 0:
		//Retrieves both the user and system times
		userTime, systemTime, err := getUserAndSystemTimes()
		if err != nil {
			return err
		}

		//Prints the retrieved info
		fmt.Printf("User Time (in ms): %s\n", userTime)
		fmt.Printf("System Time (in ms): %s\n", systemTime)
	default:
		return fmt.Errorf("%w: expected zero arguments", NoArgsNeeded)
	}

	return nil
}

func getUserAndSystemTimes() (userTime, systemTime time.Duration, err error) {
	// Gets the resource usage information for the current process
	var rusage syscall.Rusage
	err = syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	if err != nil {
		return 0, 0, fmt.Errorf("error getting resource usage information: %w", err)
	}

	//Converts the retrieved timeval ints to time.Duration format
	//time.Duration represents a duration of time in nanoseconds
	//Prints out in milliseconds because otherwise it showed in 0 seconds
	userTime = time.Duration(rusage.Utime.Sec)*time.Second + time.Duration(rusage.Utime.Usec)*time.Microsecond
	systemTime = time.Duration(rusage.Stime.Sec)*time.Second + time.Duration(rusage.Stime.Usec)*time.Microsecond

	return userTime, systemTime, nil
}
