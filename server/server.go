package server

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/spencercdixon/words/cli"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome. Head to /random to see a random word")
}

func Random(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	word := cli.RandomWord()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(word); err != nil {
		panic(err)
	}
}

func Serve() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/random", Random)

	fmt.Println("Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
