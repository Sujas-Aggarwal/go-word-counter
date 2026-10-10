package wordcounter

import (
	"fmt"
	"os"
)

// so basically what we need to do is, we can first get the list of all the files there are and then distribute the files among workers and then check the difference
// so what we can do is, first we can get the arguement of the directory from user provided and then recursively, get all the files inside to our list

func listFiles(files *[]string, path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error Reading Directory -", err)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			// then we would need to redo recursively
			listFiles(files, fmt.Sprint(path, "/", dir.Name()))
		} else {
			file_path := fmt.Sprint(path, "/", dir.Name())
			*files = append(*files, file_path)
		}
	}
}

func getWordCount(path string) int {
	content, err := os.ReadFile(path)
	if err != nil {
		println("Error Reading File Content of ", err)
	}
	return len(content)
}

func WordCount(path string) int {
	wc := 0
	files := []string{}
	listFiles(&files, path)
	for _, file := range files {
		wc += getWordCount(file)
	}
	return wc
}
