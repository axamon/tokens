package tokens

import (
	"log"
	"regexp"
	"strings"

	"github.com/axamon/lemmatizer"
	"github.com/axamon/stringset"
)

type Phrase struct {
	Sentence string
}

var vocabolarioIT *lemmatizer.Lemmatizer

func init() {
	voca, err := lemmatizer.New("it")
	if err != nil {
		log.Println(err.Error())
	}

	vocabolarioIT = voca

}

var alphabet = regexp.MustCompile(`[^a-zA-Z]`)

func (phrase Phrase) Tokenize() (tokens *stringset.StringSet) {

	words := strings.Split(phrase.Sentence, " ")

	appoggio := *stringset.NewStringSet()

	for _, word := range words {
		word = alphabet.ReplaceAllLiteralString(word, "")
		lemma := vocabolarioIT.Lemma(word)
		appoggio.Add(lemma)
	}

	tokens = &appoggio

	return tokens
}
