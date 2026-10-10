package main

import datacreator "github.com/Sujas-Aggarwal/go-word-counter/internals"

var config datacreator.Config = datacreator.Config{
	MAX_DEPTH:            3,
	MAX_FILES_PER_FOLDER: 10,
	MAX_LINES_PER_FILE:   20,
	FILE_TYPES:           []string{"py", "txt", "md", "js", "rust"},
	PATH:                 "./demo-directory"}

func main() {

	datacreator.Generate(config)
}
