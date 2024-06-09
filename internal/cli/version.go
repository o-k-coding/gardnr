package cli

import (
	"log"

	"github.com/spf13/cobra"
	version "okcoding.com/gardnr/internal/version"
)

func newVerisonCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gardnr",
		Long:  `All software has versoins. This is croku's.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			version := version.GetGardnrVersion()
			printGardnrVerison(version)
			return nil
		},
	}
}

// Pretty complex function right?
// Kept like this in case I want to add any other functionality/tooling to this
func printGardnrVerison(version string) {
	log.Println(version)
}
