package cmd

import "github.com/spf13/cobra"

var checkOffCmd = &cobra.Command{
	Use:     "check-off",
	Aliases: []string{"chkoff", "done"},
	Short:   "Check off an item.",
	Long:    "Check off an item. goes by number. the argument is the item you want to check off",
	Args:    cobra.ExactArgs(1),
	Run:     func(cmd *cobra.Command, args []string) { CheckOffItem(args[0]) },
}

func init() {
	rootCmd.AddCommand(checkOffCmd)
}
