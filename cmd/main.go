package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/dolly"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: dolly <file.vhs>")
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	cmds, err := dolly.Parse(string(b))
	if err != nil {
		fmt.Println(err)
	}
	dolly.Run(cmds)
}
