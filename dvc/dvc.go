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

	"github.com/fatih/color"
)

type DVC struct {
	dotPath        string
	dotFileName    string
	stagedFileName string
	hashes         []string
}

func New() *DVC {

	dvc := DVC{dotFileName: "toy_git.txt", stagedFileName: ".staged.txt"}
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
		_, err := file_io.CreateFile(d.dotPath)

		if err != nil {
			log.Println("Error initializing toy_git")
			log.Fatalln(err)
		}
	}

	log.Println("toy_git already Initiaized")

}

func (d *DVC) initCheck() {

	if !file_io.Exists(d.dotPath) {
		log.Fatalln("Initiaize toy_git first.")
	}

}

func (d *DVC) getUntrackFiles() []string {

	untrackedFiles := []string{}
	trackedFiles := d.getTrackedFiles()
	stagedFiles := d.getStagedFiles()

	currentDirFiles, err := file_io.GetFilesFrom(".")

	if err != nil {
		log.Fatalln("Error getting files from current dir!!")
	}

	if len(trackedFiles) == 0 && len(stagedFiles) == 0 {
		return currentDirFiles
	}

	for _, currentFile := range currentDirFiles {
		if !util.StringContains(trackedFiles, currentFile) &&
			!util.StringContains(stagedFiles, currentFile) {
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

func (d *DVC) showUntrackedFiles() {
	untrackedFiles := d.getUntrackFiles()

	if len(untrackedFiles) == 0 {
		fmt.Println("There no untracked files.")
		return
	}

	log.Println("Listing all the untracked files.")
	for _, untrackedFile := range untrackedFiles {
		color.Red(untrackedFile)
	}

}

func (d *DVC) showStagedFiles() {
	stagedFiles := d.getStagedFiles()

	if len(stagedFiles) == 0 {
		fmt.Println("There no staged files.")
		return
	}

	log.Println("Listing all the staged files.")
	for _, stagedFile := range stagedFiles {
		color.Green(stagedFile)
	}

}

func (d *DVC) StatusCommand() {
	d.initCheck()
	d.showUntrackedFiles()
	d.showStagedFiles()

}

func (d *DVC) addStagedFiles(filesToStage []string, stagedFiles []string, untrackedFiles []string) []string {

	for _, fileName := range filesToStage {
		if !util.StringContains(untrackedFiles, fileName) {
			log.Printf("%s is not in the untracked list", fileName)
			continue
		}
		stagedFiles = append(stagedFiles, fileName)
	}

	return stagedFiles

}

func (d *DVC) AddCommand(commandOptions []string) {
	d.initCheck()
	stagedFiles := d.getStagedFiles()
	untrackedFiles := d.getUntrackFiles()
	log.Println("Total untracked files : ", len(untrackedFiles))
	log.Println("Total staged files : ", len(stagedFiles))

	if commandOptions[1] == "." {
		stagedFiles = append(stagedFiles, untrackedFiles...)
	} else {
		stagedFiles = d.addStagedFiles(commandOptions[1:], stagedFiles, untrackedFiles)
	}

	file_io.WriteLinesTo(d.stagedFileName, stagedFiles)
	color.Green("Total file added : %d", len(stagedFiles))

}

func (d *DVC) getTrackedFiles() []string {

	dotFileData := d.getDotFileData()

	trackedFiles := []string{}

	for trackedFile := range dotFileData {
		trackedFiles = append(trackedFiles, trackedFile)
	}

	return trackedFiles
}

func (d *DVC) getStagedFiles() []string {

	data, err := file_io.ReadLines(d.stagedFileName)

	if err != nil {
		log.Println("Error getting the staged files.")
		log.Println(err)
		return []string{}
	}

	return data
}
