package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func cmdExtension(c *cli.Context) error {
	receiverName := c.String("name")
	results := []string{}
	err := filepath.Walk("./", checkExtensionFunctions(receiverName, &results))
	if err != nil {
		return err
	}
	for _, name := range results {
		fmt.Println("* " + name)
	}
	return nil
}

func checkExtensionFunctions(name string, results *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if info.IsDir() || !strings.Contains(info.Name(), ".kt") {
			return nil
		}
		f, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		if name == "ALL" {
			name = `\w*`
		}
		regex := `fun(\s|(\s<(\s|\w|,|:)*>\s))` + name + `(<(\s|\w|,|<|>)*>)?\.`

		matched, err := regexp.Match(regex, f)
		if err != nil {
			return err
		}
		if matched {
			*results = append(*results, info.Name())
		}
		return nil
	}
}
