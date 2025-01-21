package cli

import (
	"flag"
	"fmt"
	"os"
)

const (
	Version = "0.0.2"
	Author  = "elliot40404<avishek40404@gmail.com>"
	Name    = "easycron"
	Desc    = "Easycron is a cross platform cli app that helps configure cron jobs"
	Example = "easycron <options> <expression>"
)

func help() {
	fmt.Printf(`%s %s

%s

Usage:
  %s

Options:
  -h, --help            Show this help message
  -v, --version         Show version information
  -i                    Specify number of iterations for non-interactive mode

Examples:
  %s
`, Name, Version, Desc, Example, Example)
}

type ParsedArgs struct {
	Expr string
	Iter int
}

func ParseArgs() (ParsedArgs, error) {
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("v", false, "Show version information")
	versionLongFlag := flag.Bool("version", false, "Show version information")
	iter := flag.Int("i", 3, "num of iterations to display")

	flag.Usage = help

	flag.Parse()

	if *helpFlag {
		help()
		os.Exit(0)
	}

	if *versionFlag || *versionLongFlag {
		fmt.Printf("%s %s\n", Name, Version)
		os.Exit(0)
	}

	expr := ""

	if flag.NArg() > 0 {
		expr = flag.Arg(0)
	}

	return ParsedArgs{
		Expr: expr,
		Iter: *iter,
	}, nil
}
