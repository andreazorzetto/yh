package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
)

func printKeyValue(l yamlLine) {
	fmt.Printf("%v: %v\n", BrightRed(l.key), Yellow(l.value))
}

func printKeyNumberOrIP(l yamlLine) {
	fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))
}

func printKeyBool(l yamlLine) {
	fmt.Printf("%v: %v\n", BrightRed(l.key), Blue(l.value))
}

func printComment(l yamlLine) {
	fmt.Printf("%v %v\n", Gray(13, l.key), Gray(13, l.value))
}

func printListElement(l yamlLine) {
	fmt.Printf("%v\n", Yellow(l.raw))
}

func printInvalidLine(l yamlLine) {
	fmt.Printf("%v\n", Black(l.raw).BgBrightRed())
}

func printMultiline(l yamlLine) {
	fmt.Printf("%v\n", Gray(20-1, l.raw))
}

func printUrl(l yamlLine) {
	fmt.Printf("%v\n", Yellow(l.raw))
}
