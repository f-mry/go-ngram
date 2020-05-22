package main

import (
	"fmt"

	"github.com/negentrpy/ngram"
)

func main() {
    word := "The quick brown fox jumps over the lazy dog"
    ngram := ngram.MakeNgram(word)
    fmt.Println(ngram)
}


