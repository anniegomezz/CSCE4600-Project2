package builtins

import (
	"errors"
	"fmt"
	"os"
)

var (
	OnlyOneorZero = errors.New("Incorrect usage. Must provide either 0 or 1 arguments.")
)

func LsComm(args ...string) error {
	switch len(args) {
	case 0:
		//Calls the listCurrDir function to print the list of files and directories in the current directory
		listCurrDir()
	case 1:
		//If one argument is passed, then it'll call the listDir function
		//and print the list of files and directories from the specified directory (passed as the argument)
		listDir(args[0])
	default:
		return OnlyOneorZero //if incorrect amount of arguments passed, then print error message
	}

	return nil
}

func listCurrDir() {
	listDir(".")
}

func listDir(directory string) {
	file, err := os.Open(directory)
	if err != nil {
		fmt.Printf("Error opening the directory. Try again. %s: %v\n", directory, err) //if there was an error opening the directory, print this message
		return
	}
	defer file.Close() //close the file

	files, err := file.Readdir(0)
	if err != nil {
		fmt.Printf("Error listing directory. Try again. %s: %v\n", directory, err) //if there was an error printing the list
		return
	}

	//if no error occurred, then print the files and directories inside of the named/specified directory
	fmt.Printf("Listing files inside the directory named: %s\n", directory)
	for _, fileInfo := range files {
		fmt.Println(fileInfo.Name())
	}
	fmt.Printf("If the name of the directory is ., then you are printing all of the files inside of the current directory.\n")
	fmt.Printf("Otherwise, the files listed above are from the directory you have specified.\n")
}
