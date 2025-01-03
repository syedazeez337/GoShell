package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(args []string) {
	if len(args) == 0 {
		return
	}

	// create a command
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
	}
}
