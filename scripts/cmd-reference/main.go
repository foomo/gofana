package main

import (
	"log"
	"os"

	"github.com/foomo/gofana/cmd/gofana/command"
	"github.com/spf13/cobra/doc"
)

func main() {
	outputDir := "./docs/reference/cli"
	if len(os.Args) > 1 {
		outputDir = os.Args[1]
	}

	if err := os.MkdirAll(outputDir, 0755); err != nil {
		log.Fatal(err)
	}

	c := command.NewRoot()
	c.DisableAutoGenTag = true

	err := doc.GenMarkdownTree(c, outputDir)
	if err != nil {
		log.Fatal(err)
	}
}
