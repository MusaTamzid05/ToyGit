package main

import (
	"fmt"
	"toy_git/dvc"
)

func main() {

	dvc := dvc.New()
	for key, hash := range dvc.GetCurrentDirHashes() {
		fmt.Printf("%s => %s\n", key, hash)
	}
}
