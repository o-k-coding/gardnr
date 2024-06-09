package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	gardnr "okcoding.com/gardnr/internal/gardnr"
)

var (
	rootCmd = &cobra.Command{
		Use:   "gardnr",
		Short: "gardnr will tend your grdn",
		Long:  "Use for curating your digital garden",
	}
)

func Execute(ctx context.Context) {
	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// Define flags, configuration and commands here.
func init() {
	if err := gardnr.Gardnr.Init(context.Background()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Add Commands
	rootCmd.AddCommand(
		newVerisonCmd(),
		newCreateCmd(),
		newCountCmd(),
		newTranslateCmd(), // TODO add support to pass google creds
	)

	// Set a custom validator for this one
	rootCmd.PersistentFlags().StringVarP(&gardnr.Gardnr.Config.RootPath, "root", "r", gardnr.Gardnr.Config.RootPath, "Root path for gardnr to manage, also can use env variable GARDNR_ROOT_PATH")
	rootCmd.PersistentFlags().StringVarP(&gardnr.Gardnr.Config.TemplatePath, "templates", "t", gardnr.Gardnr.Config.TemplatePath, "Root path for gardnr template files, also can use env variable GARDNR_TEMPLATE_PATH")

	if gardnr.Gardnr.Config.RootPath == "" {
		rootCmd.MarkPersistentFlagRequired("root")
	}
}
