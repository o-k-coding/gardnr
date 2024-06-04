package cli

import (
	"path/filepath"

	"github.com/spf13/cobra"
	create "okcoding.com/grdnr/internal/create"
	grdnr "okcoding.com/grdnr/internal/grdnr"
)

func newCreateGardenPostCmd() *cobra.Command {
	var notePath string
	var postTemplate string
	var postPath string
	var title string
	var description string
	var postName string

	createDailyNoteCmd := &cobra.Command{
		Use:   "garden-post",
		Short: "Create new garden post",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			notePath := filepath.Join(grdnr.Grdnr.Config.RootPath, notePath)
			templatePath := filepath.Join(grdnr.Grdnr.Config.RootPath, grdnr.Grdnr.Config.TemplatePath, postTemplate)
			postPath := filepath.Join(grdnr.Grdnr.Config.GardenRepoPath, grdnr.Grdnr.Config.GardenRepoContentPath, postPath)
			dailyNote := create.CreateGardenPost(notePath, templatePath, postPath, postName, title, description, grdnr.Grdnr.Storage)
			if err := dailyNote.CreateFromTemplate(); err != nil {
				return err
			}
			return nil
		},
	}

	createDailyNoteCmd.Flags().StringVar(&notePath, "note", "", "Note file to turn into a garden post")
	createDailyNoteCmd.Flags().StringVar(&postTemplate, "post-template", "garden-post.mdx.tmpl", "Post template file name found in the grdnr template path (default=garden-post.mdx.tmpl)")
	createDailyNoteCmd.Flags().StringVar(&postPath, "post-path", "", "Path of where to put the post in the garden repo relative to the content directory set by GRDNR_GARDEN_REPO_CONTENT_PATH")
	createDailyNoteCmd.Flags().StringVar(&postName, "post-name", "", "Name of the post file")
	createDailyNoteCmd.Flags().StringVar(&title, "title", "", "Title of the post")
	createDailyNoteCmd.Flags().StringVar(&description, "description", "", "Description of the post")

	return createDailyNoteCmd
}
