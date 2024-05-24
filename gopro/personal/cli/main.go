	// Package cli contains a command-line interface for this code
	package main

	import (
		"flag"
		"fmt"
		"io/ioutil"
		"os"
	)

	// main is the entry point of the Go program.
	func main() {


	// It sets the flag.Usage function to provide usage information when the -h or --help flag is provided.
	// It parses the command-line flags.	
	flag.Usage = func() {
			fmt.Fprintf(os.Stderr, "Usage of %s: [list|<name>]\n\n"+
				"list: lists all files in the current directory\n"+
				"<name>: greets the specified name, using -name flag\n",
				os.Args[0])
			flag.PrintDefaults()
		}
		
		flag.Parse()
		// If the first argument is "list", it lists all files in the current directory.
		
		if flag.Arg(0) == "list" {
			files, _ := ioutil.ReadDir(".")
			for _, file := range files {
				fmt.Println(file.Name())
			}
			// If the first argument is empty, it greets the name specified by the -name flag.
		} else if flag.Arg(0) == "" {
			fmt.Printf("Hello, %s!\n", *flag.String("name", "World", "the name to greet"))

			// Otherwise, it greets the first argument.

		} else {
			fmt.Printf("Hello, %s!\n", flag.Arg(0))
		}

	}
