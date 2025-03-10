package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "CLIst",
	Short: "a CLI tool for making todo lists",
	Long:  "a CLI tool for making todo lists to increase productivity and encourage terminal use.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. an error has occurred.\n")
		os.Exit(1)
	}
}
