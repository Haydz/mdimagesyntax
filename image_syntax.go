/*
create a text file with the syntax of including images in markdown format
so that a user can copy paste the syntact into their blog

steps:
References
https://yourbasic.org/golang/list-files-in-directory/

*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func OSCheck() ([]uint8, string) {
	var osDirFormat string
	var directoryFind []uint8
	if runtime.GOOS == "windows" {
		fmt.Println(runtime.GOOS)
		//command =
		osDirFormat = "\\"
		// err := nil
		directoryFind, _ = exec.Command("cmd", "/C", "echo %cd%").Output()

	} else if runtime.GOOS == "linux" {
		fmt.Println("linux")
		directoryFind, _ = exec.Command("pwd").Output()
		osDirFormat = "/"

	}
	return directoryFind, osDirFormat

}

func main() {
	// var command string
	var osDirFormat string    // Windows = "\" Linux "/"
	var directoryFind []uint8 // saving command output into a variable

	//Checking if Windows or Linux
	directoryFind, osDirFormat = OSCheck()
	//turning cmd command back into string
	fullImageDirectory := string(directoryFind[:])

	//Splitting long directory by OS Dir structure ("\" or "/")
	fullImageDirectorySplit := strings.Split(fullImageDirectory, osDirFormat)

	//identifying current working Dir
	workingDir := fullImageDirectorySplit[len(fullImageDirectorySplit)-1]
	workingDir = strings.TrimSpace(workingDir) + osDirFormat
	fmt.Println("Working Directory is:", workingDir)
	workingDir = strings.TrimSpace(workingDir)

	//searching current directory for image files
	filesList, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	//collecting current list of files
	var filesSlice []string
	for _, files := range filesList {
		// identifying image files to apply Markdown syntax too. Place within separate slice
		if files.Name() != "image_syntax.go" && files.Name() != "test.txt" && files.Name() != "image_syntax.exe" && files.Name() != ".git" && files.Name() != "README.md" {
			filesSlice = append(filesSlice, files.Name())
		}
	}

	//opening file to write output to
	outputFile, err := os.Create("mdsyntax_output.txt")
	if err != nil {
		fmt.Println(err)

	}
	defer outputFile.Close()
	//iterating through images files
	for i := range filesSlice {
		pathForFile := workingDir + filesSlice[i]
		//writing image markdown syanx for iamges
		stringToWrite := "![" + filesSlice[i] + "](" + pathForFile + ")"
		fmt.Println(stringToWrite)
		_, err := io.WriteString(outputFile, stringToWrite+"\n")
		if err != nil {
			fmt.Println(err)
		}
	}

}
