package main

import (
	"fmt"
	"os"

	"github.com/gold-chen-five/ggit/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Eroor %v\n", err)
		os.Exit(1)
	}
}
