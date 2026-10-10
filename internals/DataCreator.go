package datacreator

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type Config struct {
	MAX_DEPTH            int
	MAX_FILES_PER_FOLDER int
	MAX_LINES_PER_FILE   int
	FILE_TYPES           []string
	PATH                 string
}

var DICT = []string{
	"Apple A round, crisp fruit that is red, green, or yellow.",
	"Brave Ready to face and endure danger or pain; showing courage.",
	"Cautious Being careful to avoid potential problems or dangers.",
	"Delight A high degree of pleasure or joy.",
	"Empathy The ability to understand and share the feelings of another.",
	"Friendly Kind and pleasant; acting like a friend.",
	"Gratitude The feeling of being thankful for something.",
	"Honest Telling the truth or free of deceit; fair and sincere.",
	"Immaculate Perfectly clean, neat, or tidy.",
	"Joyful Feeling, expressing, or causing great pleasure and happiness.",
	"Kindness The quality of being friendly, generous, and considerate.",
	"Lucid Clear and easy to understand.",
	"Mitigate To make something less severe, serious, or painful.",
	"Nostalgia A sentimental longing or affection for the past.",
	"Optimist A person who tends to be hopeful and confident about the future.",
	"Pragmatic Dealing with things sensibly and realistically.",
	"Quick Moving or doing something fast; prompt.",
	"Resilient Able to recover quickly from tough situations.",
	"Sincere Free from pretense or deceit; honest and genuine.",
}

var incrementalNumber int = 0

func getNumber() int {
	defer func() { incrementalNumber += 1 }()
	return incrementalNumber
}

func createDirectory(path string) string {
	dir_name := fmt.Sprint(path, "/directory_", getNumber())
	err := os.Mkdir(dir_name, 0777)
	if err != nil {
		fmt.Println("Error Creating Directory", err)
		os.Exit(1)
	}
	return dir_name
}

func getRandomContent(length int) string {
	randomSize := rand.IntN(length)
	var content strings.Builder
	for range randomSize {
		content.WriteString(DICT[rand.IntN(len(DICT))])
	}
	return content.String()
}

func createFile(size int, path string, format string) string {
	file_name := fmt.Sprint(path, "/file_", getNumber(), ".", format)
	err := os.WriteFile(file_name, []byte(getRandomContent(size)), 0777)
	if err != nil {
		fmt.Println("Error Creating File")
		os.Exit(1)
	}
	return file_name
}

func Generate(config Config) {
	// lets finally write the actual logic of dummy data creation now.
	os.RemoveAll(config.PATH) // we will empty it first and then fill it up like a good boy
	// now we will go inside each level and actually make directories or files inside it randomly
	// let's first create directories, later we can populate them with files
	dirs := []string{}
	current_path := config.PATH
	err := os.Mkdir(config.PATH, 0777)
	if err != nil {
		fmt.Println("Error Emptying Path Dir")
	}
	for range config.MAX_DEPTH {
		// we will keep creating dirs inside dir
		dir := createDirectory(current_path)
		current_path = dir
		dirs = append(dirs, dir)
	}
	// we got our dirs setup now
	// now lets populate them with files
	for _, dir := range dirs {
		// get total number of files to create
		files := rand.IntN(config.MAX_FILES_PER_FOLDER)
		for range files {
			createFile(config.MAX_LINES_PER_FILE, dir, config.FILE_TYPES[rand.IntN(len(config.FILE_TYPES))])
		}
	}
	// and we are done
}
