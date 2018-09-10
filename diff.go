package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func diffE(_ *cobra.Command, args []string) error {
	if _, err := os.Stat("go.mod"); err != nil {
		return err
	}

	// Check if the vendor directory exists and it has a .git directory.
	if _, err := os.Stat("vendor"); err != nil {
		if os.IsNotExist(err) {
			return errors.New("vendor directory is missing; please run `vendiff init` to vendor the dependencies")
		}
		return err
	}
	if _, err := os.Stat("vendor/.git"); err != nil {
		if os.IsNotExist(err) {
			return errors.New("vendor directory exists but is not being tracked, aborting")
		}
		return err
	}

	cmd := exec.Command("git", "diff")
	if len(args) > 0 {
		cmd.Args = append(cmd.Args, "--")
		cmd.Args = append(cmd.Args, args...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = "vendor"
	return cmd.Run()
}
