package cli

import (
	"log"

	"github.com/spf13/cobra"
	"okcoding.com/grdnr/internal/count"
	grdnr "okcoding.com/grdnr/internal/grdnr"
)

func newCountWordsCmd() *cobra.Command {

	var filePath string

	createDailyNoteCmd := &cobra.Command{
		Use:   "words",
		Short: "Count words in file",
		Long:  "Counts words separated by whitespace in file",
		RunE: func(cmd *cobra.Command, args []string) error {
			count, err := count.CountWords(grdnr.Grdnr.RootPath, filePath)
			if err != nil {
				return err
			}
			log.Printf("word count for %s: %d", filePath, count)
			return nil
		},
	}

	createDailyNoteCmd.Flags().StringVarP(&filePath, "file", "f", "", "The path of the file to count words from")

	return createDailyNoteCmd
}
