package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/logrusorgru/aurora"
)

const version = "0.2.1"

func main() {

	// checking the args, someone out there might need help
	checkArgs(os.Args)

	// get the juice
	scanner := bufio.NewScanner(os.Stdin)

	var text []string

	for scanner.Scan() {
		if scanner.Text() == "EOF" {
			break
		}
		text = append(text, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	nextLineMaybeAComment := false
	indentBeforeComment := 0

	// parse the juice
	for _, line := range text {
		// Read lines of pasted YAML

		if (nextLineMaybeAComment == true) && (countIndentSpaces(line) > indentBeforeComment) {
			// Found multiline comment or configmap, not treated as YAML at all

			fmt.Printf("%v\n", Gray(20-1, line))

		} else if strings.Contains(line, ":") {
			// Line contains ":"

			lineSlice := strings.Split(line, ":")

			k := lineSlice[0]
			v := strings.TrimSpace(strings.Join(lineSlice[1:len(lineSlice)], ":"))

			_, err := strconv.Atoi(strings.ReplaceAll(v, ".", ""))

			if string(strings.TrimSpace(k)[0]) == "#" {
				// Line is a comment
				fmt.Printf("%v %v\n", Gray(13, k), Gray(13, v))

			} else if err == nil {
				// Value is a number
				fmt.Printf("%v: %v\n", BrightRed(k), Blue(v))

			} else if (strings.ToLower(v) == "true") || (strings.ToLower(v) == "false") {
				// the value is boolean
				fmt.Printf("%v: %v\n", BrightRed(k), Blue(v))

			} else {
				// Value is a word
				fmt.Printf("%v: %v\n", BrightRed(k), Yellow(v))
			}

			if containsChompingIndicator(v) {
				// Found possible multiline comment or configmap
				// If this check is validated with the next line the text is highlighted as multiline comment
				nextLineMaybeAComment = true
				indentBeforeComment = countIndentSpaces(line)

			} else {
				nextLineMaybeAComment = false
			}

		} else if len(strings.TrimSpace(line)) > 0 {
			// Line doesn't contain ":" and it's not an empty line

			if string(strings.TrimSpace(line)[0]) == "#" {
				// Line is a comment
				fmt.Printf("%v\n", Gray(13, line))

			} else if string(strings.TrimSpace(line)[0]) == "-" {
				// Line is an item of a list
				fmt.Printf("%v\n", Yellow(line))

			} else {
				// Line is not valid
				fmt.Printf("%v\n", Black(line).BgBrightRed())
			}

			nextLineMaybeAComment = false

		} else {
			// Empty or spaces only line
			fmt.Println(line)
		}

	}

}

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

func countIndentSpaces(s string) int {
	// This function checks how many indentation spaces where used
	// before chomping indicator to catch a possible multiline comment or config
	count := 0

	for _, v := range s {
		if string(v) == " " {
			count += 1
		} else {
			break
		}
	}
	return count
}

func containsChompingIndicator(s string) bool {
	// this function checks for multline chomping indicator
	indicators := []string{">", ">-", ">+", "|", "|-", "|+"}

	for _, v := range indicators {
		if strings.Contains(s, v) {
			return true
		}
	}
	return false
}
