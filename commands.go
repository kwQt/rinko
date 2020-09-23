package main

import "github.com/urfave/cli/v2"

var commands = []*cli.Command{
	{
		Name:   "comment",
		Usage:  "check whether link comments are written above the definition of particular classes",
		Action: cmdComment,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "name", Value: "Fragment", Usage: "specify suffix of file name"},
			&cli.StringFlag{Name: "extension", Value: "kt", Aliases: []string{"ext"}, Usage: "specify file extension"},
			&cli.BoolFlag{Name: "all", Value: false, Aliases: []string{"a"}, Usage: "display all results"},
		},
	},
}
