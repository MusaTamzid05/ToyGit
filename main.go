package main

import (
	"fmt"
	"log"
	"toy_git/file_io"
)

func main() {

	lines, err := file_io.ReadLines("./test.txt")

	if err != nil {
		log.Fatalln(err)
	}

	for i, line := range lines {
		fmt.Printf("%d => %s\n", i, line)
	}
}
