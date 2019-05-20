package dvc // distribute vision control

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"toy_git/file_io"
)

type DVC struct {
	dotPath     string
	dotFileName string
	hashes      []string
}

func New() *DVC {

	dvc := DVC{dotFileName: "toy_git.txt"}
	dvc.init()
	return &dvc
}

func (d *DVC) init() {
	currentDir, err := os.Getwd()

	if err != nil {
		log.Println(err)
	}

	d.dotPath = currentDir + string(os.PathSeparator) + d.dotFileName
}

func (d *DVC) GetCurrentDirHashes() map[string]string {

	hashes := make(map[string]string)

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		currentHash := md5.New()
		fmt.Fprintf(currentHash, "%v", info.IsDir())
		fmt.Fprintf(currentHash, "%v", info.ModTime())
		fmt.Fprintf(currentHash, "%v", info.Mode())
		fmt.Fprintf(currentHash, "%v", info.Name())
		fmt.Fprintf(currentHash, "%v", info.Size())

		hashes[info.Name()] = fmt.Sprintf("%x", currentHash.Sum(nil))

		return nil
	})

	if err != nil {
		fmt.Println("Error : ", err)
	}

	return hashes

}

func (d *DVC) InitCommand() {

	if !file_io.Exists(d.dotPath) {
		file_io.CreateFile(d.dotPath)
	}

}
