package cmd

import "github.com/spf13/cobra"

var addItemCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{},
	Short:   "add item to the plan file",
	Long:    "add item to the plan file",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		AddItem(args[0])
	},
}

func init() {
	rootCmd.AddCommand(addItemCmd)
}
