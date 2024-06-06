package main

import (
	"fmt"
	"os"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: failed getting working directory: %v", os.Args[0], err)
		os.Exit(1)
	}

	c := clipboard.New()
	err = c.CopyText(path)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: failed copying working directory to clipboard: %v", os.Args[0], err)
		os.Exit(1)
	}

	os.Exit(0)
}
