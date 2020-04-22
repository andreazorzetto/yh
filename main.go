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

        // Read lines of pasted YAML
	for _, line := range text {

		// Line contains ":"
		if strings.Contains(line, ":") {

			lineSlice := strings.Split(line, ":")

			k := lineSlice[0]
			v := strings.TrimSpace(strings.Join(lineSlice[1:len(lineSlice)],":"))

			_, err := strconv.Atoi(strings.ReplaceAll(v, ".", ""))

			if string(strings.TrimSpace(k)[0]) == "#" {
				// Line is a comment
				fmt.Printf("%v %v\n", Gray(13, k), Gray(13, v))

			} else if err == nil {
				// Value is a number
				fmt.Printf("%v: %v\n", BrightRed(k), Cyan(v))

			} else {
				// Value is a word
				fmt.Printf("%v: %v\n", BrightRed(k), Yellow(v))

			}
			// Line doesn't contain ":"
		} else {
			// Line is a comment
			if string(strings.TrimSpace(line)[0]) == "#" {
				fmt.Printf("%v\n", Gray(13, line))

				// Line is an item of a list
			} else if string(strings.TrimSpace(line)[0]) == "-" {
				fmt.Printf("%v\n", Yellow(line))
				// Line is not valid
			} else {
				fmt.Printf("%v\n", Black(line).BgBrightRed())
			}
		}

	}

}
