package file_io

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {

	filePtr, err := os.Open(path)
	lines := []string{}

	if err != nil {
		return lines, err
	}

	defer filePtr.Close()

	scanner := bufio.NewScanner(filePtr)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil

}

func Exists(path string) bool {

	fmt.Println(path)
	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func CreateFile(path string) error {
	fmt.Println("Creating ", path)

	_, err := os.Create(path)

	if err != nil {
		return err
	}

	return nil

}
