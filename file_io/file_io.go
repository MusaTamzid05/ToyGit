package file_io

import (
	"bufio"
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
