package main

import (
	"github.com/elliot40404/easycron/internal/parser"
	"github.com/elliot40404/easycron/internal/renderer"
)

func main() {
	cronParser := parser.NewCronParser()
	renderer.TviewRenderer(cronParser)
}
