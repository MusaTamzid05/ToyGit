package main

import (
	"toy_git/dvc"
)

func main() {

	dvc := dvc.New()
	dvc.InitCommand()
	//dvc.StatusCommand()
	dvc.AddCommand([]string{"add", "."})
	dvc.StatusCommand()
}
