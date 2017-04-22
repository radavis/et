package common

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// String prompt.
func String(prompt string, args ...interface{}) string {
	fmt.Printf(prompt, args...)
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(s, "\r\n")
}

// StringRequired prompt (required).
func StringRequired(prompt string, args ...interface{}) string {
	s := String(prompt, args...)

	if strings.Trim(s, " ") == "" {
		return StringRequired(prompt)
	}

	return s
}

// Confirm continues prompting until the input is boolean-ish.
func Confirm(prompt string, args ...interface{}) bool {
	s := String(prompt, args...)
	switch s {
	case "yes", "y", "Y":
		return true
	case "no", "n", "N":
		return false
	default:
		return Confirm(prompt, args...)
	}
}
