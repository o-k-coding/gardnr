package create

import (
	"bufio"
	"log"
	"os"
	"strings"

	"okcoding.com/grdnr/internal/template"
)

type GardenPostTemplateData struct {
	Title       string
	Description string
}

// TODO allow to pass in template data - but also
// n+ will be to generate the template data using an llm and ask for confirmation in the CLI
// ALSO in the cli build out a way to preview the mdx
type GardenPost struct {
	NotePath     string
	TemplatePath string
	PostPath     string
	PostName     string
	GardenPostTemplateData
}

func CreateGardenPost(notePath string, templatePath string, postPath string, postName string, title string, description string) *GardenPost {
	return &GardenPost{
		NotePath:     notePath,
		TemplatePath: templatePath,
		PostPath:     postPath,
		PostName:     postName,
		GardenPostTemplateData: GardenPostTemplateData{
			Title:       title,
			Description: description,
		},
	}
}

func (g *GardenPost) CreateFromTemplate() error {
	log.Println("reading data from", g.NotePath)
	noteFile, err := os.Open(g.NotePath)

	// Next read the file... line by line? something. and feed it to the md parser to build a structure
	if err != nil {
		return err
	}

	fileScanner := bufio.NewScanner(noteFile)
	fileScanner.Split(bufio.ScanLines)

	var file *os.File

	for fileScanner.Scan() {
		t := fileScanner.Text()
		if file == nil && len(t) != 0 && t[0] == '#' {
			// TODO function to parse lines into DS - tree based with headers as nodes
			if g.GardenPostTemplateData.Title == "" {
				g.GardenPostTemplateData.Title = strings.Trim(t[1:], " ")
			}
			if g.PostName == "" {
				templateParts := strings.Split(g.TemplatePath, ".")
				fileExt := templateParts[len(templateParts)-2]
				// TODO this assumes the format of name.ext.tmpl and will break in many cases
				g.PostName = strings.ToLower(strings.ReplaceAll(g.GardenPostTemplateData.Title, " ", "-")) + "." + fileExt
			}
			file, err = template.CreateFileFromTemplate(g.TemplatePath, g.PostPath, g.PostName, g.GardenPostTemplateData)
			if err != nil {
				log.Printf("Error creating post file %s from template %s : %e", g.PostPath, g.TemplatePath, err)
				return err
			}
			defer file.Close()
			// Don't include the top level header
			continue
		}
		if file == nil {
			continue
		}
		if _, err := file.WriteString(fileScanner.Text() + "\n"); err != nil {
			log.Printf("Error writing line from %s to %s : %e", g.NotePath, g.PostPath, err)
			return err
		}
	}

	return nil
}

func (g *GardenPost) SaveDraft() error {
	// TODO create a branch and commit the file and push
	return nil
}

func (g *GardenPost) Publish() error {
	// TODO commit and push
	return nil
}
