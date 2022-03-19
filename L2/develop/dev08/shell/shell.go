package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Shell structure
type Shell struct {
	commands []string
}

// NewShell shell's structure constructor
func NewShell() *Shell {
	return &Shell{
		commands: make([]string, 0),
	}
}

// Run starts shell
func (s *Shell) Run() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("$ ")
		if scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())

			switch line {
			case "":
				continue
			case "exit":
				os.Exit(0)
			default:
				if err := s.ExecuteCommands(line); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}
		}
	}
}

// ExecuteCommands runs commands
func (s *Shell) ExecuteCommands(line string) error {
	s.commands = strings.Split(line, " | ")

	for _, command := range s.commands {
		args := strings.Fields(command)

		switch args[0] {
		case "cd":
			if len(args) < 2 {
				home, _ := os.UserHomeDir()
				os.Chdir(home)

				return nil
			}

			if err := os.Chdir(args[1]); err != nil {
				return fmt.Errorf("-shell: cd: %s: no such file or directory", args[1])
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout

			if err := cmd.Run(); err != nil {
				return fmt.Errorf("%s: command not found", cmd.String())
			}
		}
	}

	return nil
}
