package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Word struct {
	PartOfSpeech string `json:"type"`
	Definition   string `json:"defenition"`
	Example      string `json:"example"`
	Self         string `json:"self,omitempty"`
}

type WordResults []Word

func (w Word) Display() {
	fmt.Println("Part of speech: ", w.PartOfSpeech)
	fmt.Println("Definition: ", w.Definition)
	fmt.Println("Example Sentence: ", w.Example)
}

// Convenience function for displaying definitions of saved words
func (w Word) DisplayDef() {
	fmt.Printf("%s (%s): %s\n", w.Self, w.PartOfSpeech, w.Definition)
}

func Fetch(word string) WordResults {
	url := fmt.Sprintf("https://owlbot.info/api/v1/dictionary/%s", word)
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		log.Fatalf("Requesting the word %s failed.", word)
	}

	results := WordResults{}
	json.NewDecoder(res.Body).Decode(&results)

	// Save what word was searched in the word
	for i := 0; i < len(results); i++ {
		results[i].Self = word
	}

	return results
}

// selectWord shows users all the different definitions results and prompts them
// to select the definition that works best for them.
func SelectWord(words WordResults) (int, Word) {
	for i, w := range words {
		// Add one to index so results start at 1 and not 0
		fmt.Printf("Result #%d\n", i+1)
		w.Display()
		fmt.Println()
	}
	selection := PromptInt("Which result do you want to use? ")
	return selection, words[selection-1]
}

// check if this is first word being saved
// unmarshal the text and decode it into types
// add the new word and marshal it and save it to the file
func SaveWord(word Word) {
	words := getWords()
	words = append(words, word)
	writeWords(words)
}

func ListWords() {
	savedWords := getWords()
	for _, word := range savedWords {
		word.DisplayDef()
	}
}

func RandomWord() {
	words := getWords()
	rand.Seed(time.Now().UTC().UnixNano())
	i := rand.Intn(len(words))
	words[i].DisplayDef()
}

func savedWordsPath() string {
	home, _ := homedir.Dir()
	return home + "/.words/saved.json"
}

func wordsDirPath() string {
	home, _ := homedir.Dir()
	return home + "/.words"
}

func getWords() WordResults {
	confirmLocalStorageExists()

	raw, err := ioutil.ReadFile(savedWordsPath())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var results WordResults
	json.Unmarshal(raw, &results)
	return results
}

func writeWords(words WordResults) {
	confirmLocalStorageExists()

	bs, err := json.Marshal(words)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(savedWordsPath(), bs, os.ModeAppend); err != nil {
		log.Fatal(err)
	}
}

// Determign whether or not a file path exists
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

func confirmLocalStorageExists() {
	// confirm our .words directory exists
	if _, err := os.Stat(wordsDirPath()); os.IsNotExist(err) {
		os.Mkdir(wordsDirPath(), 0755)
	}

	// confirm we have a saved.json file to use for writing saved words
	path := savedWordsPath()
	doesExist, err := exists(path)
	if !doesExist || err != nil {
		answer := Confirm("No words found.  Want to start saving words?")
		if answer {
			var data []byte
			ioutil.WriteFile(path, data, 0755)
		} else {
			os.Exit(1)
		}
	}
}
