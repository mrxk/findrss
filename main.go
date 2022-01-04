package main

import (
	"os"

	"github.com/mrxk/findrss/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
