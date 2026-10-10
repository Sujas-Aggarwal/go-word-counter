package main

import (
	"fmt"

	wordcounter "github.com/Sujas-Aggarwal/go-word-counter/internals/WordCounter"
)

func main() {
	fmt.Println("Total Word Count -", wordcounter.WordCount("./demo-directory"))
}
