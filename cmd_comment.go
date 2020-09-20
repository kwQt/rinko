package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/urfave/cli/v2"
)

func cmdComment(c *cli.Context) error {

	results := map[string]string{}

	suffix := c.String("name")

	err := filepath.Walk("./", checkComments(suffix, results))
	if err != nil {
		return err
	}

	showAll := c.Bool("all")
	existFiles := map[string]string{}
	emptyFiles := map[string]string{}

	for name, url := range results {
		if url == "" {
			emptyFiles[name] = url
		} else {
			existFiles[name] = url
		}
	}

	for name := range emptyFiles {
		fmt.Println("* " + name)
	}

	if showAll {
		for name, url := range existFiles {
			fmt.Println("* " + name + ": " + url)
		}
	}

	return nil
}

func checkComments(suffix string, results map[string]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.Contains(info.Name(), suffix+".") {
			return nil
		}

		f, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		singleLineCommentRegex := `//(.|\s)*class(.|\s)*` + suffix
		multipleLineCommentRegex := `/\*(.|\s)*\*/(.|\s)*class(.|\s)*` + suffix
		linkRegex := `https?://.*`

		rS, err := regexp.Compile(singleLineCommentRegex)
		if err != nil {
			return err
		}
		rM, err := regexp.Compile(multipleLineCommentRegex)
		if err != nil {
			return err
		}
		rlink, err := regexp.Compile(linkRegex)
		if err != nil {
			return err
		}

		data := rS.Find(f)
		url := rlink.Find(data)
		if url != nil {
			results[info.Name()] = string(url)
			return nil
		}

		data = rM.Find(f)
		url = rlink.Find(data)

		results[info.Name()] = string(url)

		return nil
	}
}
