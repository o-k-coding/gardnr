package cli

import (
	"path/filepath"

	"github.com/spf13/cobra"
	create "okcoding.com/gardnr/internal/create"
	gardnr "okcoding.com/gardnr/internal/gardnr"
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
			notePath := filepath.Join(gardnr.Gardnr.Config.RootPath, notePath)
			templatePath := filepath.Join(gardnr.Gardnr.Config.RootPath, gardnr.Gardnr.Config.TemplatePath, postTemplate)
			postPath := filepath.Join(gardnr.Gardnr.Config.GardenRepoPath, gardnr.Gardnr.Config.GardenRepoContentPath, postPath)
			dailyNote := create.CreateGardenPost(notePath, templatePath, postPath, postName, title, description, gardnr.Gardnr.Storage)
			if err := dailyNote.CreateFromTemplate(); err != nil {
				return err
			}
			return nil
		},
	}

	createDailyNoteCmd.Flags().StringVar(&notePath, "note", "", "Note file to turn into a garden post")
	createDailyNoteCmd.Flags().StringVar(&postTemplate, "post-template", "garden-post.mdx.tmpl", "Post template file name found in the gardnr template path (default=garden-post.mdx.tmpl)")
	createDailyNoteCmd.Flags().StringVar(&postPath, "post-path", "", "Path of where to put the post in the garden repo relative to the content directory set by GARDNR_GARDEN_REPO_CONTENT_PATH")
	createDailyNoteCmd.Flags().StringVar(&postName, "post-name", "", "Name of the post file")
	createDailyNoteCmd.Flags().StringVar(&title, "title", "", "Title of the post")
	createDailyNoteCmd.Flags().StringVar(&description, "description", "", "Description of the post")

	return createDailyNoteCmd
}
