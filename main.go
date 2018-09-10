package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "vendiff",
	Run: func(cmd *cobra.Command, args []string) {
		if err := diffE(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s.\n", err)
			os.Exit(1)
		}
	},
}

var InitCommand = &cobra.Command{
	Use: "init",
	Run: func(cmd *cobra.Command, args []string) {
		if err := initE(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s.\n", err)
			os.Exit(1)
		}
	},
}

var CleanCommand = &cobra.Command{
	Use: "clean",
	Run: func(cmd *cobra.Command, args []string) {
		if err := cleanE(cmd, args); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s.\n", err)
			os.Exit(1)
		}
	},
}

func main() {
	if err := Command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s.\n", err)
		os.Exit(1)
	}
}

func init() {
	Command.AddCommand(InitCommand)
	Command.AddCommand(CleanCommand)
	CleanCommand.Flags().BoolVarP(&cleanFlags.Force, "force", "f", false, "force the removal of the vendor directory even if there are changes")
}
