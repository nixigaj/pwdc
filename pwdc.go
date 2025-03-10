package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/tiagomelo/go-clipboard/clipboard"
)

func copyUsingOSC52(text string) error {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	osc52 := fmt.Sprintf("\x1b]52;0;%s\x07", encoded)
	_, err := os.Stdout.WriteString(osc52)
	return err
}

func main() {
	path, err := os.Getwd()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s: failed getting working directory: %v\n", os.Args[0], err)
		os.Exit(1)
	}

	if os.Getenv("PWDC_USE_OSC52") == "1" {
		if err := copyUsingOSC52(path); err != nil {
			fmt.Fprintf(os.Stderr, "%s: failed copying working directory via OSC52: %v\n", os.Args[0], err)
			os.Exit(1)
		}
		return
	}

	c := clipboard.New()
	err = c.CopyText(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: failed copying working directory to clipboard, trying OSC52: %v\n", os.Args[0], err)
		if err := copyUsingOSC52(path); err != nil {
			fmt.Fprintf(os.Stderr, "%s: failed copying working directory via OSC52: %v\n", os.Args[0], err)
			os.Exit(1)
		}
	}
}
