package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andreazorzetto/yh/highlight"
)

const version = "0.4.0"

func main() {
	log.SetFlags(0) // disable timestamp of log package

	// Checking the args, someone out there might need help
	checkArgs(os.Args)

	h, err := highlight.Highlight(os.Stdin)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Print(h)
}

// Check args if passed
// Show help
func checkArgs(a []string) {
	if len(a) >= 2 {
		// Someone's looking for...

		if a[1] == "version" {
			// version
			fmt.Println(version)
			os.Exit(0)

		} else if a[1] == "help" {
			// help
			fmt.Println("You don't really need to read this! \nJust pipe me some YAML. I don't bite")
			fmt.Println("\nExample:")
			fmt.Println("\tkubectl get myNastyPod -o yaml | yh")
			fmt.Println("\nCommands:")
			fmt.Println("\thelp: get this helpful help")
			fmt.Println("\tversion: get the version")
			os.Exit(0)
		} else {
			// trolling
			fmt.Println("Not really sure of what you want! Maybe try help or version.")
			os.Exit(0)
		}
	}
}
