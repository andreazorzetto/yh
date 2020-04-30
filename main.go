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

type line struct {
	line  string
	key   string
	value string
}

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
	for _, read := range text {
		// Read lines of pasted YAML

		l := line{
			line: read,
		}

		if (nextLineMaybeAComment == true) && (countIndentSpaces(read) > indentBeforeComment) {
			// Found multiline comment or configmap, not treated as YAML at all

			fmt.Printf("%v\n", Gray(20-1, l.line))

		} else if l.isKeyValue() {

			l.getKeyValue()

			if isComment(l.key) {
				fmt.Printf("%v %v\n", Gray(13, l.key), Gray(13, l.value))

			} else if isNumberOrIP(l.value) {
				fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))

			} else if isBoolean(l.value) {
				fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))

			} else {
				// Value is a word
				fmt.Printf("%v: %v\n", BrightRed(l.key), Yellow(l.value))
			}

			if containsChompingIndicator(l.value) {
				// Found possible multiline comment or configmap
				// If this check is validated with the next line the text is highlighted as multiline comment
				nextLineMaybeAComment = true
				indentBeforeComment = countIndentSpaces(l.line)

			} else {
				nextLineMaybeAComment = false
			}

		} else if len(strings.TrimSpace(l.line)) > 0 {
			// Line doesn't contain ":" and it's not an empty line

			if string(strings.TrimSpace(l.line)[0]) == "#" {
				// Line is a comment
				fmt.Printf("%v\n", Gray(13, l.line))

			} else if string(strings.TrimSpace(l.line)[0]) == "-" {
				// Line is an item of a list
				fmt.Printf("%v\n", Yellow(l.line))

			} else {
				// Line is not valid
				fmt.Printf("%v\n", Black(l.line).BgBrightRed())
			}

			nextLineMaybeAComment = false

		} else {
			// Empty or spaces only line
			fmt.Println(l.line)
		}

	}

}

func (l line) isKeyValue() bool {
	if strings.Contains(l.line, ":") {
		return true
	} else {
		return false
	}

}

func (l *line) getKeyValue() {
	t := strings.Split(l.line, ":")

	l.key = t[0]
	l.value = strings.TrimSpace(strings.Join(t[1:len(t)], ":"))
}

func isComment(s string) bool {
	if string(strings.TrimSpace(s)[0]) == "#" {
		// Line is a comment
		return true
	}
	return false
}

func isBoolean(s string) bool {
	if (strings.ToLower(s) == "true") || (strings.ToLower(s) == "false") {
		// Line is a boolean value
		return true
	}
	return false
}

func isNumberOrIP(s string) bool {
	_, err := strconv.Atoi(strings.ReplaceAll(s, ".", ""))
	if err == nil {
		// Line is a number or IP
		return true
	} else {
		return false
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
