package main

import (
	"fmt"
	"os"

	"github.com/open-policy-agent/conftest/internal/commands"
)

func main() {
	fmt.Println("post-transfer conftest")

	if err := commands.NewDefaultCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
