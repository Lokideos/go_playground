package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 1 {
		// Get filename
		filename := os.Args[1]
		// Open file for reading
		file, err := os.Open(filename)

		// If there is an error (e.g. there is no file with this name),
		// then print an error and exit the programm
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}

		// Read the contents of the file and write them to the terminal
		io.Copy(os.Stdout, file)
		// Close the file
		file.Close()
	} else {
		fmt.Println("No filename was provided")
	}
}
