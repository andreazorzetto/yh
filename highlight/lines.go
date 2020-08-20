package highlight

import (
	"strconv"
	"strings"
)

type yamlLine struct {
	raw   string
	key   string
	value string
}

func (l yamlLine) isKeyValue() bool {
	if strings.Contains(l.raw, ":") {
		return true
	}
	return false
}

func (l *yamlLine) getKeyValue() {
	t := strings.Split(l.raw, ":")

	l.key = t[0]
	l.value = strings.TrimSpace(strings.Join(t[1:len(t)], ":"))
}

func (l yamlLine) isComment() bool {
	if string(strings.TrimSpace(l.raw)[0]) == "#" {
		// Line is a comment
		return true
	}
	return false
}

func (l yamlLine) valueIsBoolean() bool {
	if (strings.ToLower(l.value) == "true") || (strings.ToLower(l.value) == "false") {
		// Line is a boolean value
		return true
	}
	return false
}

func (l yamlLine) valueIsNumberOrIP() bool {
	_, err := strconv.Atoi(strings.ReplaceAll(l.value, ".", ""))
	if err == nil {
		// Line is a number or IP
		return true
	}
	return false
}

func (l yamlLine) isEmptyLine() bool {
	if len(strings.TrimSpace(l.raw)) > 0 {
		return false
	}
	return true
}

func (l yamlLine) isElementOfList() bool {
	if string(strings.TrimSpace(l.raw)[0]) == "-" {
		return true
	}
	return false
}

func (l yamlLine) indentationSpaces() int {
	// This function checks how many indentation spaces where used
	// before chomping indicator to catch a possible multiline comment or config
	count := 0

	for _, v := range l.raw {
		if string(v) == " " {
			count += 1
		} else {
			break
		}
	}
	return count
}

func (l yamlLine) valueContainsChompingIndicator() bool {
	// this function checks for multline chomping indicator
	indicators := []string{">", ">-", ">+", "|", "|-", "|+"}

	for _, in := range indicators {
		if strings.Contains(l.value, in) {
			return true
		}
	}
	return false
}
