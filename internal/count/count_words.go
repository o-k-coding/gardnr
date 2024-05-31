package count

import (
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

// TODO figure out if hashing the file is faster than counting the words
// COuld leverage a cache of the file hash with word count to spare some cycles

func CountWords(rootPath string, filePath string) (int, error) {
	path := path.Join(rootPath, filePath)
	log.Println("reading data from", path)
	// Be warned this will read the whole fie into memory!
	contents, err := os.ReadFile(path)
	// Next read the file... line by line? something. and feed it to the md parser to build a structure
	if err != nil {
		return 0, err
	}

	return CleanAndCountWords(string(contents)), nil
}

func CleanAndCountWords(text string) int {
	// naively
	// 1. turn the file contents into a string
	// 2. replace newlines and punctuation with spaces
	// 3. merge any multi spaces
	// 4. split by spaces and count the length
	replacePunctuation := regexp.MustCompile(`[^\w\s]`)
	replaceWhitespace := regexp.MustCompile(`\s`)
	replaceMultiSpaces := regexp.MustCompile(`\s{2,}`)
	cleanedText := replacePunctuation.ReplaceAllString(text, "")
	cleanedText = replaceWhitespace.ReplaceAllString(cleanedText, " ")
	cleanedText = replaceMultiSpaces.ReplaceAllString(cleanedText, " ")
	cleanedText = strings.TrimSpace(cleanedText)
	wordCount := strings.Count(cleanedText, " ") + 1
	return wordCount
}
