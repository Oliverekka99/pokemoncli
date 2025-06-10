package main

import (
	"fmt"
	"os"
)

func callbackExit() error {
	fmt.Println("Exiting the REPL. Goodbye!")
	// Perform any necessary cleanup here before exiting
	// For example, saving state or closing connections
	// os.Exit(0) is not called here to allow for graceful shutdown
  //want to save state or close connections
	os.Exit(0) // Exit the program
	return nil
}