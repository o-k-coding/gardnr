package cli

import (
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	create "okcoding.com/gardnr/internal/create"
	gardnr "okcoding.com/gardnr/internal/gardnr"
)

func newCreateDailyNoteCmd() *cobra.Command {
	var addDays int
	var noteTemplate string

	createDailyNoteCmd := &cobra.Command{
		Use:   "daily-note",
		Short: "Create new daily note file",
		Long:  "By default creates a new daily note file with the correct dates and times filled out based on the template",
		RunE: func(cmd *cobra.Command, args []string) error {
			noteDate := time.Now().AddDate(0, 0, addDays)
			templatePath := filepath.Join(gardnr.Gardnr.Config.RootPath, gardnr.Gardnr.Config.TemplatePath, noteTemplate)
			dailyNote := create.CreateDailyNote(gardnr.Gardnr.Config.RootPath, templatePath, noteDate)
			if err := dailyNote.CreateFromTemplate(); err != nil {
				return err
			}
			return nil
		},
	}

	createDailyNoteCmd.Flags().IntVar(&addDays, "add-days", 0, "Create a note for today + n days. Will rewrite if it already exists. Value can be negative (default=0)")
	createDailyNoteCmd.Flags().StringVar(&noteTemplate, "note-template", "daily-note.md.tmpl", "Note template file name found in the gardnr template path (default=daily_note.md.tmpl)")

	return createDailyNoteCmd
}
