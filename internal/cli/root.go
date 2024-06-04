package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	grdnr "okcoding.com/grdnr/internal/grdnr"
)

var (
	rootCmd = &cobra.Command{
		Use:   "grdnr",
		Short: "grdnr will tend your grdn",
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
	if err := grdnr.Grdnr.Init(context.Background()); err != nil {
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
	rootCmd.PersistentFlags().StringVarP(&grdnr.Grdnr.Config.RootPath, "root", "r", grdnr.Grdnr.Config.RootPath, "Root path for grdnr to manage, also can use env variable GRDNR_ROOT_PATH")
	rootCmd.PersistentFlags().StringVarP(&grdnr.Grdnr.Config.TemplatePath, "templates", "t", grdnr.Grdnr.Config.TemplatePath, "Root path for grdnr template files, also can use env variable GRDNR_TEMPLATE_PATH")

	if grdnr.Grdnr.Config.RootPath == "" {
		rootCmd.MarkPersistentFlagRequired("root")
	}
}
