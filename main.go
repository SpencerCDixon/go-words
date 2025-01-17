package main

import (
	"flag"
	"fmt"
	"github.com/spencercdixon/words/cli"
	"github.com/spencercdixon/words/server"
	"log"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("\twords add <word>\n")
		fmt.Printf("\twords list\n")
		fmt.Printf("\twords random\n")
		fmt.Printf("\twords serve\n")
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
	case "serve":
		server.Serve()
	default:
		fmt.Printf("Unknown command %s\n\n", command)
		flag.Usage()
	}
}

func addWord(words []string) {
	if len(words) == 0 {
		log.Fatalln("words: Need a word to add")
	}

	defs := cli.Fetch(words[0])
	_, selectedDef := cli.SelectWord(defs)
	cli.SaveWord(selectedDef)
}

// Wrapper around ListWords in case I ever decide to add flags/options in the listing process
func listWords() {
	cli.ListWords()
}

// Wrapper around ListWords in case I ever decide to add flags/options in the listing process
func randomWord() {
	cli.DisplayRandomWord()
}
