package wordcounter

import (
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
)

func listFiles(files *[]string, path string) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("Error Reading Directory -", err)
		os.Exit(1)
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			listFiles(files, fmt.Sprint(path, "/", dir.Name()))
		} else {
			file_path := fmt.Sprint(path, "/", dir.Name())
			*files = append(*files, file_path)
		}
	}
}

func getWordCount(path string, results chan<- int) {
	content, err := os.ReadFile(path)
	if err != nil {
		println("Error Reading File Content of ", err)
		results <- 0
		return
	}
	results <- len(content)
}

func WordCount(path string) int {
	f, _ := os.Create("./cpu.prof")
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	wc := 0
	var wg sync.WaitGroup
	files := []string{}

	// Find all files first
	listFiles(&files, path)

	results := make(chan int, len(files))

	for _, file := range files {
		filePath := file
		wg.Go(func() {
			getWordCount(filePath, results)
		})
	}

	// This monitors the WaitGroup and closes the channel when workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Read from channel until it's closed
	for count := range results {
		wc += count
	}

	return wc
}
