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

// handle buildin commands
func HandleBuildInCommands(args []string) bool {
	if len(args) == 0 {
		return false
	}

	switch args[0] {
	case "exit":
		fmt.Println("Goodbye!")
		os.Exit(0)
	case "cd":
		if len(args) < 2 {
			fmt.Println("cd: missing argument")
			return true
		}
		err := os.Chdir(args[1])
		if err != nil {
			fmt.Println("cd error:", err)
		}
		return true
	case "pwd":
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error:", err)
			return true
		}
		fmt.Println(wd)
		return true
	case "ls":
		files, err := os.ReadDir(getDir())
		if err != nil {
			fmt.Println("Error:", err)
			return true
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}

		return true
	}

	return false
}


// get directory
func getDir() string {
	wd, _ := os.Getwd()
	return wd
}