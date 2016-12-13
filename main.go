package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("\twords add <word>\n")
		fmt.Printf("\twords list\n")
		fmt.Printf("\twords random\n")
		flag.PrintDefaults()
	}

	// Parse flag and display usage if no args are given
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(0)
	}

	command := os.Args[1]
	args := os.Args[2:]

	switch command {
	case "add":
		addWord(args)
	case "list":
		listWords()
	case "random":
		randomWord()
	default:
		fmt.Printf("Unknown command %s\n\n", command)
		flag.Usage()
	}
}

func addWord(words []string) {
	if len(words) == 0 {
		log.Fatalln("words: Need a word to add")
	}

	defs := Fetch(words[0])
	_, selectedDef := SelectWord(defs)
	SaveWord(selectedDef)
}

// Wrapper around ListWords in case I ever decide to add flags/options in the
// listing process
func listWords() {
	ListWords()
}

func randomWord() {
	RandomWord()
}
