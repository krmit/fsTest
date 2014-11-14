package main

import (
	"flag"
	"fmt"
	"github.com/dustin/randbo"
	"io"
	"os"
	"strconv"
)

var smallLetters = "abcdefghijklmnopqrstuvwxyz"
var bigLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

type fsTestOption struct {
	maxFile  int
	sizeFile int
	maxDir   int
	maxDeep  int
}

func main() {
	//Parse commadline flags.
	option := &fsTestOption{}

	flag.IntVar(&option.maxFile, "f", 5, "Number of files to create in each dir.")
	flag.IntVar(&option.maxDeep, "d", 5, "Number of subdirectories to be created")
	flag.IntVar(&option.sizeFile, "s", 5, "The size of the files to generate")
	flag.IntVar(&option.maxDir, "D", 5, "Number of directories on a level.")
	flag.Parse()

	option.sizeFile = option.sizeFile * 1024
	os.Mkdir("tmpTest", 0777)
	os.Chdir("tmpTest")
	numberOfFiles := createDir(*option, option.maxDeep)
	os.Chdir("..")
	fmt.Println("Number of files: " + strconv.Itoa(numberOfFiles))
}

func createDir(option fsTestOption, level int) int {
	for i := 0; i < option.maxFile; i++ {
		_ = createRandomFile(option, i)
	}

	numberOfFiles := option.maxFile
	if level > 0 {
		for i := 0; i < option.maxDir; i++ {
			dirName := string(smallLetters[i])
			os.Mkdir(dirName, 0777)
			_ = os.Chdir(dirName)
			level--
			numberOfFiles += createDir(option, level)
			_ = os.Chdir("..")
		}
	}
	return numberOfFiles
}

func createRandomFile(option fsTestOption, fileNumber int) error {
	r := randbo.New()
	fileName := string(bigLetters[fileNumber]) + ".blob"
	f, _ := os.Create(fileName)
	io.CopyN(f, r, int64(option.sizeFile))
	return nil
}
