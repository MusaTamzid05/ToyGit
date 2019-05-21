package file_io

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
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

	if _, err := os.Stat(path); err == nil {
		return true
	}

	return false
}

func CreateFile(path string) (*os.File, error) {
	fmt.Println("Creating ", path)

	fPtr, err := os.Create(path)

	if err != nil {
		return nil, err
	}

	return fPtr, nil

}

func GetFilesFrom(path string) ([]string, error) {

	files := []string{}

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, info.Name())

		return nil
	})

	if err != nil {
		return files, err
	}

	return files, nil

}

func WriteLinesTo(path string, lines []string) {

	fPtr, err := CreateFile(path)

	if err != nil {
		log.Fatalln(err)
	}

	defer fPtr.Close()

	for _, line := range lines {
		fmt.Fprintln(fPtr, line)
	}
}
