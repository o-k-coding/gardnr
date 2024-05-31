package cli

import "github.com/spf13/cobra"

func newCountCmd() *cobra.Command {
	countCommand := &cobra.Command{
		Use:   "count",
		Short: "Command for counting things",
		Long:  "Support for counting things in the framework of note taking",
	}
	countCommand.AddCommand(
		newCountWordsCmd(),
	)
	return countCommand
}
