package create_test

import (
	"testing"
	"time"

	create "okcoding.com/grdnr/internal/create"
)

func TestCreateDailyNote(t *testing.T) {
	rootPath := "/test/path"
	templatePath := "/test/path/.grdnr/templates/daily-note.tmpl"
	currentTime, err := time.Parse("2006-01-02 15:04", "2022-04-24 09:00")

	if err != nil {
		t.Fatal("Error creating currentTime for note data, sad")
	}

	dailyNote := create.CreateDailyNote(rootPath, templatePath, currentTime)

	expNoteName := "04-24-2022.md"

	if dailyNote.NoteName != expNoteName {
		t.Errorf("dailyNote.NoteName = %s; want %s", dailyNote.NoteName, expNoteName)
	}

	expNoteDir := "/test/path/daily/2022/april"

	if dailyNote.NoteDir != expNoteDir {
		t.Errorf("dailyNote.NoteDir = %s; want %s", dailyNote.NoteDir, expNoteDir)
	}

	expNotePath := "/test/path/daily/2022/april/04-24-2022.md"

	if dailyNote.NotePath != expNotePath {
		t.Errorf("dailyNote.NotePath = %s; want %s", dailyNote.NotePath, expNotePath)
	}

	expCurrentDate := "04-24-2022"

	if dailyNote.DailyNoteTemplateData.CurrentDate != expCurrentDate {
		t.Errorf("dailyNote.TemplateData.CurrentDate = %s; want %s", dailyNote.DailyNoteTemplateData.CurrentDate, expCurrentDate)
	}

	expStartingTime := "09:00"

	if dailyNote.DailyNoteTemplateData.StartingTime != expStartingTime {
		t.Errorf("dailyNote.TemplateData.StartingTime = %s; want %s", dailyNote.DailyNoteTemplateData.StartingTime, expStartingTime)
	}
}
