package create

import (
	"bufio"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"okcoding.com/grdnr/internal/objectstorage"
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
	Storage      objectstorage.ObjectStorage
	NotePath     string
	TemplatePath string
	PostPath     string
	PostName     string
	GardenPostTemplateData
}

func CreateGardenPost(notePath string, templatePath string, postPath string, postName string, title string, description string, storage objectstorage.ObjectStorage) *GardenPost {
	return &GardenPost{
		Storage:      storage,
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
	imageImport := false

	if g.GardenPostTemplateData.Title == "" {
		noteFileName := filepath.Base(g.NotePath)
		noteFileExtension := filepath.Ext(noteFileName)
		g.GardenPostTemplateData.Title = strings.Trim(noteFileName[0:len(noteFileName)-len(noteFileExtension)], " ")
	}
	// TODO this all seems very hacky and not very flexible
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

	for fileScanner.Scan() {
		t := fileScanner.Text()
		if len(t) != 0 && t[0] == '#' {
			// TODO function to parse lines into DS - tree based with headers as nodes

			// Don't include the top level header
			continue
		}
		if file == nil {
			continue
		}
		text := fileScanner.Text()

		if strings.Contains(text, "![") {
			// TODO upload image file to storage
			// TODO replace with image tag
			text, err = handleImage(context.Background(), filepath.Dir(g.NotePath), text, g.Storage)
			if err != nil {
				log.Printf("Error handling image in %s : %e", g.NotePath, err)
				return err
			}
			if !imageImport {
				file.WriteString("import { Image } from 'astro:assets';\n\n")
				imageImport = true
			}
		}
		if _, err := file.WriteString(text + "\n"); err != nil {
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

func handleImage(ctx context.Context, noteDir string, text string, storage objectstorage.ObjectStorage) (string, error) {
	// TODO this only handles obsidian style image tags
	start := -1
	imagePath := ""
	// TODO skip the beginning square brackets too
	for i, c := range text {
		if c == '!' && start == -1 {
			start = i + 1
			continue
		}
		if c == '[' && start != -1 {
			start++
			continue
		}
		if c == ']' && start != -1 {
			imagePath = text[start:i]
			break
		}
	}

	key := filepath.Base(imagePath)
	// TODO generate alt text for the file
	text = "<Image src='https://images.okcoding.io/" + key + "' alt='" + key + "' inferSize={true}/>"

	// now check if the file already exists in storage bucket
	exists, err := storage.CheckFileExists(ctx, key)

	if err != nil {
		return "", err
	}
	if !exists {
		// TODO check file size and type
		// TODO this always assumes the image is a sibling to the note
		file, err := os.Open(filepath.Join(noteDir, key)) // TODO this is relative - need to make it absolute
		if err != nil {
			return "", err
		}
		err = storage.UploadFile(ctx, key, "image/"+filepath.Ext(key), file, 0)
		if err != nil {
			return "", err
		}
	}

	// upload if not and replace the text with the image tag using the correct url
	return text, nil
}
