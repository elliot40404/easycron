package cli

import (
	"flag"
	"fmt"
	"os"
)

const (
	Version = "0.0.1"
	Author  = "elliot40404<avishek40404@gmail.com>"
	Name    = "easycron"
	Desc    = "Easycron is a cross platform cli app that helps configure cron jobs"
	Example = "easycron <options>"
)

func help() {
	fmt.Printf(`%s %s

%s

Usage:
  %s

Options:
  -h, --help            Show this help message
  -v, --version         Show version information

Examples:
  %s
`, Name, Version, Desc, Example, Example)
}

func ParseArgs() error {
	helpFlag := flag.Bool("help", false, "Show help message")
	versionFlag := flag.Bool("v", false, "Show version information")
	versionLongFlag := flag.Bool("version", false, "Show version information")

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

	return nil
}
