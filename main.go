package main

import (
	"bufio"
	"fmt"
	"os"

	. "github.com/logrusorgru/aurora"
)

const version = "0.3.0"

func main() {
	foundChompingIndicator := false
	indentationSpacesBeforeComment := 0

	// checking the args, someone out there might need help
	checkArgs(os.Args)

	// get the juice
	text, err := readTextFromStdin()

	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	// parse the juice
	for _, line := range text {

		l := yamlLine{
			raw: line,
		}

		if (foundChompingIndicator == true) && (l.indentationSpaces() > indentationSpacesBeforeComment) {
			// Found multiline comment or configmap, not treated as YAML at all

			fmt.Printf("%v\n", Gray(20-1, l.raw))

		} else if l.isKeyValue() {

			l.getKeyValue()

			if l.isComment() {
				fmt.Printf("%v %v\n", Gray(13, l.key), Gray(13, l.value))

			} else if l.valueIsNumberOrIP() {
				fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))

			} else if l.valueIsBoolean() {
				fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))

			} else {
				// Value is a word
				fmt.Printf("%v: %v\n", BrightRed(l.key), Yellow(l.value))
			}

			if l.valueContainsChompingIndicator() {
				// Found possible multiline comment or configmap
				// If this check is validated with the next line the text is highlighted as multiline comment

				foundChompingIndicator = true
				indentationSpacesBeforeComment = l.indentationSpaces()

			} else {
				foundChompingIndicator = false
			}

		} else if !l.isEmptyLine() {

			if l.isComment() {
				fmt.Printf("%v\n", Gray(13, l.raw))

			} else if l.isElementOfList() {
				fmt.Printf("%v\n", Yellow(l.raw))

			} else {
				// Line is not valid
				fmt.Printf("%v\n", Black(l.raw).BgBrightRed())
			}

			foundChompingIndicator = false

		} else if l.isEmptyLine() {
			// Empty or spaces only line
			fmt.Println(l.raw)
		}

	}

}

func readTextFromStdin() ([]string, error) {
	// Read all the text from Stdin and return as a []string

	scanner := bufio.NewScanner(os.Stdin)

	var text []string

	for scanner.Scan() {
		if scanner.Text() == "EOF" {
			break
		}
		text = append(text, scanner.Text())
	}

	return text, scanner.Err()
}

func checkArgs(a []string) {
	// Check args if passed
	// Show help

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
