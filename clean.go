package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var cleanFlags struct {
	Force bool
}

func cleanE(_ *cobra.Command, args []string) error {
	if _, err := os.Stat("vendor"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
		return nil
	}

	if _, err := os.Stat("vendor/.git"); err != nil && os.IsNotExist(err) {
		return errors.New("vendor directory was not created by vendiff, aborting")
	}

	if !cleanFlags.Force {
		// Verify there are no changes in the vendor directory.
		cmd := exec.Command("git", "diff", "--quiet")
		cmd.Dir = "vendor"
		if err := cmd.Run(); err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				return errors.New("changes exist in vendor; ensure that patches have been created and then use -f to remove")
			}
			return err
		}
	}
	return os.RemoveAll("vendor")
}
