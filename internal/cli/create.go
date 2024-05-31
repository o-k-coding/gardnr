package cli

import (
	"github.com/spf13/cobra"
)

func newCreateCmd() *cobra.Command {
	// Create is just a base "verb" command that needs a more specific "noun" command to tell what to create
	// The nouns will then have specific flags to apply
	// If a flag were to apply to all base commands of create, they would be added here as a
	// persistent flag
	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Command for creating files",
		Long:  "Support for creating files that fit into the workflow and framework of note taking provided by grdnr",
	}

	createCmd.AddCommand(
		newCreateDailyNoteCmd(),
		newCreateGardenPostCmd(),
	)

	return createCmd
}
