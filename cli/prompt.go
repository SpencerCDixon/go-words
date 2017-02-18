package cli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// PromptString asks users a question, collects some input
// then cleans it up and returns the string to be used
func PromptString(question string, args ...interface{}) string {
	fmt.Printf(question, args...)
	s, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	return strings.TrimRight(s, "\r\n")
}

// PromptInt asks users a question, collects some input
// and then converts it to an int and returns
func PromptInt(question string, args ...interface{}) int {
	s := PromptString(question, args...)
	n, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("words: %v", err)
	}
	return n
}

// Confirm is a quick and easy way to get true/false confirmation
// from the user.  It auto formats questions with (y/n) notation.
func Confirm(question string, args ...interface{}) bool {
	formattedQuestion := question + " (y/n) "
	s := PromptString(formattedQuestion, args...)
	switch s {
	case "yes", "y", "Y":
		return true
	case "no", "n", "N":
		return false
	default:
		return Confirm(formattedQuestion, args...)
	}
}
