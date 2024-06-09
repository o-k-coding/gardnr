package create

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"okcoding.com/gardnr/internal/template"
)

type DailyNoteTemplateData struct {
	CurrentDate  string
	StartingTime string
}

type DailyNote struct {
	RootPath              string
	NoteDate              time.Time
	CopyNoteDate          *time.Time
	NotePath              string
	NoteDir               string
	NoteName              string
	TemplatePath          string
	DailyNoteTemplateData // Data used to fill the template file TODO make this more dynamic?
}

// This is a calculation
// Note there is a subtlety here that is very important
// By passing in the time, this can be a calculation because that is just data.
// If we called time.Now() in this function, this would become an action
// because then it would always depend on WHEN it was called. which would be much harder to test
// This is true for rootPath as well, since ultimately that data needs to be loaded in.
func CreateDailyNote(rootPath string, templatePath string, noteDate time.Time) *DailyNote {
	currentDateFmt := noteDate.Format("01-02-2006")
	startTime := noteDate.Format("15:04")
	templateData := DailyNoteTemplateData{
		CurrentDate:  currentDateFmt,
		StartingTime: startTime,
	}

	noteName := getDailyNoteName(noteDate)
	noteDir := getDailyNoteDir(rootPath, noteDate)

	return &DailyNote{
		RootPath: rootPath,
		NoteDate: noteDate,
		NotePath: filepath.Join(noteDir, noteName),
		NoteDir:  noteDir,
		NoteName: noteName,
		// TemplateName: "daily-note.md",
		TemplatePath:          templatePath,
		DailyNoteTemplateData: templateData,
	}
}

// This is an action because it depends on the file structure, which could change at any time
func (dailyNote *DailyNote) CreateFromTemplate() error {
	file, err := template.CreateFileFromTemplate(dailyNote.TemplatePath, dailyNote.NoteDir, dailyNote.NoteName, dailyNote.DailyNoteTemplateData)
	if err != nil {
		log.Printf("Error create note file %s from template %s : %e", dailyNote.NotePath, dailyNote.TemplatePath, err)
		return err
	}

	// TODO rewrite to allow copying from any date
	if dailyNote.CopyNoteDate != nil {
		previousNotePath := getMostRecentDayNotePath(dailyNote.RootPath, dailyNote.NoteDate)
		log.Println("copy previous day note", previousNotePath)
		// TODO next read the previous note and write data to the new note
		// TODO create data structure for daily note file
		data, err := readDataFromNote(previousNotePath)
		if err != nil {
			return err
		}
		writeDateToNote(file, data)
	}

	log.Printf("Created new daily note file %s", dailyNote.NotePath)
	return nil
}

// Action that checks to see what is the most recent note prior to the one passed in
func getMostRecentDayNotePath(rootPath string, date time.Time) string {
	// TODO want some form of short circuit if there are no previous days
	previousDate := date.AddDate(0, 0, -1)
	previousNoteDir := getDailyNoteDir(rootPath, previousDate)
	previousNoteName := getDailyNoteDir(rootPath, previousDate)
	previousNotePath := fmt.Sprintf("%s/%s", previousNoteDir, previousNoteName)

	if _, err := os.Stat(previousNotePath); err == nil {
		return previousNotePath
	} else if errors.Is(err, os.ErrNotExist) {
		log.Println("Could not find note at", previousNotePath)
		return getMostRecentDayNotePath(rootPath, previousDate)
	}
	// Should not hit here...
	return ""
}

// Calculation to generate a daily note directory given a date and a rootpath
func getDailyNoteDir(rootPath string, date time.Time) string {
	year, month, _ := date.Date()
	return filepath.Join(rootPath, "daily", strconv.Itoa(year), strings.ToLower(month.String()))
}

// Calculation to generate a daily note file name given a date
func getDailyNoteName(date time.Time) string {
	dateFmt := date.Format("01-02-2006")
	return dateFmt + ".md"
}

// Action that reads daily note data from the given note path
func readDataFromNote(notePath string) (interface{}, error) {
	log.Println("reading data from", notePath)
	_, err := os.ReadFile(notePath)
	// Next read the file... line by line? something. and feed it to the md parser to build a structure
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Action that writes daily note data to the given file TODO
func writeDateToNote(file *os.File, data interface{}) {
	log.Println("writing data to", file.Name())
}

// TODO probably a calculation that takes in file data (bytes slices?) and packs it into data structure
