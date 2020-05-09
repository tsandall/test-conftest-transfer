package main

import (
	"fmt"
	"os"

	"github.com/tsandall/test-conftest-transfer/internal/commands"
)

func main() {
	fmt.Println("post-transfer conftest")

	if err := commands.NewDefaultCommand().Execute(); err != nil {
		os.Exit(1)
	}
}
