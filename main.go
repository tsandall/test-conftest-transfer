package main

import (
	"fmt"
	"os"

	"github.com/tsandall/test-conftest-transfer/internal/commands"
)

func main() {
	fmt.Println("post-transfer conftest (2)")

	if err := commands.NewDefaultCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
