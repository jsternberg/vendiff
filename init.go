package main

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func initE(_ *cobra.Command, args []string) error {
	if _, err := os.Stat("go.mod"); err != nil {
		return err
	}

	// If the vendor directory does not exist, create it using `go mod vendor`.
	if _, err := os.Stat("vendor"); err != nil {
		cmd := exec.Command("go", "mod", "vendor")
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GO111MODULE=on")
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}

	// Initialize the vendor directory as a git repository if it doesn't have a .git directory.
	if _, err := os.Stat("vendor/.git"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		cmd := exec.Command("git", "init")
		cmd.Stderr = os.Stderr
		cmd.Dir = "vendor"
		if err := cmd.Run(); err != nil {
			return err
		}

		cmd = exec.Command("git", "add", "-A")
		cmd.Stderr = os.Stderr
		cmd.Dir = "vendor"
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
