package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "rinko"
	app.Usage = "A trivial source code checker (especially for Android application development)."
	app.Commands = commands

	app.Run(os.Args)
}
