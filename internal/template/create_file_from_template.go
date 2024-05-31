package template

import (
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func CreateFileFromTemplate(templatePath string, destinationDir string, destinationFile string, templateData interface{}) (*os.File, error) {
	t_node, err := template.ParseFiles(templatePath)

	if err != nil {
		log.Printf("Error parsing template file at %s : %e", templatePath, err)
		return nil, err
	}

	// Check that the folder exists first
	_, err = os.Stat(destinationDir)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(destinationDir, 0755) // TODO look at the perm value
		if errDir != nil {
			log.Printf("Error creating dir at %s : %e", destinationDir, err)
			return nil, errDir
		}
	}

	filePath := filepath.Join(destinationDir, destinationFile)
	file, err := os.Create(filePath)
	if err != nil {
		log.Printf("Error creating file at %s : %e", filePath, err)
		return nil, err
	}

	err = t_node.Execute(file, templateData)
	if err != nil {
		log.Printf("Error executing template %s on file %s : %e", templatePath, filePath, err)
		return nil, err
	}

	log.Printf("Created new file %s from template %s", filePath, templatePath)
	return file, nil
}
