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

		// Line doesn't contains ":"
		if strings.Contains(line, ":") {

			k := strings.Split(line, ":")[0]
			v := strings.TrimSpace(strings.Split(line, ":")[1])

			_, err := strconv.Atoi(strings.ReplaceAll(v, ".", ""))

			// Line is a comment
			if string(strings.TrimSpace(k)[0]) == "#" {
				fmt.Printf("%v %v\n", Gray(13, k), Gray(13, v))

				// Value is a number
			} else if err == nil {
				fmt.Printf("%v: %v\n", BrightRed(k), Cyan(v))

				// Value is a word
			} else {
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
