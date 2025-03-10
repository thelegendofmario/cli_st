package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createPlanCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"new"},
	Short:   "Create new plan.txt file",
	Long:    "Create new plan.txt file. WARNING: THIS WILL OVERWRITE YOUR OLD PLAN.TXT FILE!!",
	Args:    cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		loc := CreatePlan()
		if loc != "fail" {
			fmt.Printf("created blank plan file at %s\n", CreatePlan())
		} else {
			fmt.Printf("File not created! Error!\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(createPlanCmd)
}
