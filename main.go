package main

import (
	"bufio"
	"fmt"
	"os"
)

const version = "0.3.0"

func main() {
	// Service vars
	foundChompingIndicator := false
	indentationSpacesBeforeComment := 0

	// Checking the args, someone out there might need help
	checkArgs(os.Args)

	// Warm-up the engine
	scanner := bufio.NewScanner(os.Stdin)

	// Get the juice
	for scanner.Scan() {

		if scanner.Text() == "EOF" {
			break
		}

		// Check for errors during Stdin read
		err := scanner.Err()

		if err != nil {
			println("Error:", err)
			os.Exit(1)
		}

		// Drink the juice
		l := yamlLine{raw: scanner.Text()}

		if (foundChompingIndicator == true) && (l.indentationSpaces() > indentationSpacesBeforeComment) {
			// Found multiline comment or configmap, not treated as YAML at all

			printMultiline(l)

		} else if l.isKeyValue() {
			// This is a valid YAML key: value line

			// Extract key and value in their own vars in the struct
			l.getKeyValue()

			if l.isComment() {
				// This line is a comment

				printComment(l)

			} else if l.valueIsNumberOrIP() {
				// The value is a number or an IP address x.x.x.x

				printKeyNumberOrIP(l)

			} else if l.valueIsBoolean() {
				// The value is boolean true or false

				printKeyBool(l)

			} else {
				// The is a normal key/value line

				printKeyValue(l)
			}

			if l.valueContainsChompingIndicator() {
				// This line contains a chomping indicator, sign of a possible multiline text coming next

				// Setting flag for next execution
				foundChompingIndicator = true

				// Saving current number of indentation spaces
				indentationSpacesBeforeComment = l.indentationSpaces()

			} else {
				// Resetting multiline flag
				foundChompingIndicator = false
			}

		} else if !l.isEmptyLine() {
			// This is not a YAML key: value line and is not empty

			if l.isComment() {
				// This line is a comment

				printComment(l)

			} else if l.isElementOfList() {
				// This line is an element of a list

				printListElement(l)

			} else {
				// This line is not valid YAML

				printInvalidLine(l)
			}

			foundChompingIndicator = false

		} else if l.isEmptyLine() {
			// This is an empty line

			fmt.Println(l.raw)
		}

	}

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
