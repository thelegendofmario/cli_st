package cmd

import "github.com/spf13/cobra"

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete an item. (takes a number!)",
	Long:    "Delete an item. goes by number. the argument is the item you want to delete",
	Args:    cobra.ExactArgs(1),
	Run:     func(cmd *cobra.Command, args []string) { DeleteItem(args[0]) },
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
