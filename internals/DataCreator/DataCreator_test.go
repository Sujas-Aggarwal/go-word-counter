package datacreator

import (
	"os"
	"testing"
)

var config Config = Config{
	MAX_DEPTH:            3,
	MAX_FILES_PER_FOLDER: 10,
	MAX_LINES_PER_FILE:   20,
	FILE_TYPES:           []string{"py", "txt", "md", "js", "go", "rust"},
	PATH:                 "."}

func TestCreateDirectory(t *testing.T) {
	dir := createDirectory(config.PATH)
	err := os.Remove(dir)
	if err != nil {
		t.Errorf("Failed - Directory Creation")
	}
}

func TestCreateFile(t *testing.T) {
	file := createFile(config.MAX_LINES_PER_FILE, config.PATH, "py")
	_, err1 := os.ReadFile(file)
	if err1 != nil {
		t.Errorf("Failed - File could not be read")
	}
	err := os.Remove(file)
	if err != nil {
		t.Errorf("Failed - File Creation")
	}
}
