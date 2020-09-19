package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var inappropriateFileNames []string
var existUrls = map[string]string{}

func main() {
	err := filepath.Walk("./", checkComments)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nInappropriate files")
	for _, name := range inappropriateFileNames {
		fmt.Println("* " + name)
	}

	fmt.Println("\nAlready satisfied files")
	for name, url := range existUrls {
		fmt.Println("* " + name + ": " + url)
	}
	fmt.Println("")
}

func checkComments(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	if !strings.Contains(info.Name(), "Fragment") {
		return nil
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	rdoc, err := regexp.Compile(`/\*(.|\s)*\*/(.|\s)*class(.|\s)*Fragment`)
	if err != nil {
		return err
	}
	data := rdoc.Find(f)

	rlink, err := regexp.Compile(`https?://.*`)
	if err != nil {
		return err
	}
	url := rlink.Find(data)

	if url == nil {
		inappropriateFileNames = append(inappropriateFileNames, info.Name())
		return nil
	}

	existUrls[info.Name()] = string(url)

	return nil
}
