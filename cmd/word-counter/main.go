package main

import (
	"fmt"
	"os"

	wordcounter "github.com/Sujas-Aggarwal/go-word-counter/internals/WordCounter"
)

func main() {
	args := os.Args
	path := "./demo-directory"
	if len(args) > 1 {
		path = args[1]
	}
	fmt.Println("Total Word Count -", wordcounter.WordCount(path))
}
