package main

import (
	"log/slog"
	"os"

	"github.com/elliot40404/easycron/internal/cli"
	"github.com/elliot40404/easycron/internal/parser"
	"github.com/elliot40404/easycron/internal/renderer"
)

func main() {
	err := cli.ParseArgs()
	if err != nil {
		slog.Error("something went wrong", "error", err)
		os.Exit(1)
	}
	cronParser := parser.NewCronParser()
	renderer.TviewRenderer(cronParser)
}
