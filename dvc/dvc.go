package dvc // distribute vision control

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"toy_git/file_io"
	"toy_git/util"
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

func (d *DVC) initCheck() {

	if !file_io.Exists(d.dotPath) {
		log.Fatalln("Initiaize toy_git first.")
	}

}

func (d *DVC) getUntrackFiles() []string {

	untrackedFiles := []string{}
	trackedFiles := d.getTrackedFiles()

	if len(trackedFiles) == 0 {
		return untrackedFiles
	}

	currentDirFiles, err := file_io.GetFilesFrom(".")

	if err != nil {
		log.Fatalln("Error getting files from current dir!!")
	}

	for _, currentFile := range currentDirFiles {
		if !util.StringContains(trackedFiles, currentFile) {
			untrackedFiles = append(untrackedFiles, currentFile)
		}

	}

	return untrackedFiles

}

func (d *DVC) getDotFileData() map[string]string {

	dotFileData := make(map[string]string)
	lines, err := file_io.ReadLines(d.dotPath)

	if err != nil {
		log.Println("Error getting untracked files.")
		log.Fatalln(err)
	}

	if len(lines) == 0 {
		return dotFileData
	}

	for _, line := range lines {
		data := strings.Split(line, " ")
		dotFileData[data[0]] = data[1]
	}

	return dotFileData

}

func (d *DVC) StatusCommand() {
	d.initCheck()
	untrackedFiles := d.getUntrackFiles()

	if len(untrackedFiles) == 0 {
		fmt.Println("There no untracked files.")
	}

	fmt.Println("Listing all the untracked files.")
	for _, currentFile := range untrackedFiles {
		fmt.Println(currentFile)
	}
}

func (d *DVC) getTrackedFiles() []string {

	dotFileData := d.getDotFileData()

	trackedFiles := []string{}

	for trackedFile := range dotFileData {
		trackedFiles = append(trackedFiles, trackedFile)
	}

	return trackedFiles
}
