package main

import (
	"fmt"
	"os"

	"github.com/axamon/tokens"
)

const messaggiobase string = "i  Presentatori sono lavoratori che si Presentano in tv, e la tv è il male assoluto. La tv è malefica."

func main() {
	var frase tokens.Phrase
	var frase2 tokens.Phrase

	frase.Sentence = os.Args[1]
	fmt.Println(frase.Sentence)
	frase2.Sentence = os.Args[2]
	fmt.Println(frase2.Sentence)

	tokens1 := frase.Tokenize()
	tokens2 := frase2.Tokenize()

	fmt.Printf("Parole inserite %d, %d\n", tokens1.Len(), tokens2.Len())
	listatokens := tokens1.Strings()
	listatokens2 := tokens2.Strings()
	fmt.Println("tokens ", listatokens)
	fmt.Println("tokens2 ", listatokens2)

	intersecation := tokens1.Intersect(tokens2)

	fmt.Println(intersecation.Strings())
}
