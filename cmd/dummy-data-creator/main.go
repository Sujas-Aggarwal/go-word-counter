package main

import (
	"os"
	"strconv"

	datacreator "github.com/Sujas-Aggarwal/go-word-counter/internals/DataCreator"
)

var config datacreator.Config = datacreator.Config{
	MAX_DEPTH:            5,
	MAX_FILES_PER_FOLDER: 1000,
	MAX_LINES_PER_FILE:   20,
	FILE_TYPES:           []string{"py", "txt", "md", "js", "rust"},
	PATH:                 "./demo-directory"}

func main() {
	args := os.Args
	if len(args) > 1 {
		config.MAX_DEPTH, _ = strconv.Atoi(args[1])
	}
	datacreator.Generate(config)
}
