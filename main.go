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

func main() {
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

	for _, line := range text {
		// Read lines of pasted YAML

		if strings.Contains(line, ":") {
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

		} else {
			// Empty or spaces only line
			fmt.Println(line)
		}

	}

}
