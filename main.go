package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		printHead() // [username@hostname workingDirectory]$

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		if len(args) < 2 {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}
			return os.Chdir(homeDir)
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func printHead() {
	// I'm not doing anything with the errors which could be dangerous
	hostName, _ := os.Hostname()
	user, _ := user.Current()
	wd, _ := os.Getwd()
	sepratedDirs := strings.Split(wd, "/")
	currentDir := sepratedDirs[len(sepratedDirs)-1]
	fmt.Printf("[%s@%s %s]", user.Username, hostName, currentDir)
	fmt.Print("$ ")
}
