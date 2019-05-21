package main

import (
	"toy_git/dvc"
)

func main() {

	dvc := dvc.New()
	dvc.InitCommand()
	dvc.AddCommand([]string{"add", "dvc.go"})
	dvc.StatusCommand()
}
