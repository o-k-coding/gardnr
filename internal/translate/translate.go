package translate

import (
	"context"
	"fmt"

	translate "cloud.google.com/go/translate/apiv3"
	translatepb "google.golang.org/genproto/googleapis/cloud/translate/v3"
)

// TODO should keep track of translates already made, and avoid calling the API twice for the same translations

func TranslateText(sourceLang string, targetLang string, text string) (string, error) {
	ctx := context.Background()
	// TODO improvement to have a long lived client?
	client, err := translate.NewTranslationClient(ctx)
	if err != nil {
		return "", fmt.Errorf("NewTranslationClient: %v", err)
	}
	defer client.Close()

	req := &translatepb.TranslateTextRequest{
		Parent:             fmt.Sprintf("projects/%s/locations/global", "gardnr"),
		SourceLanguageCode: sourceLang,
		TargetLanguageCode: targetLang,
		MimeType:           "text/plain", // Mime types: "text/plain", "text/html"
		Contents:           []string{text},
	}

	resp, err := client.TranslateText(ctx, req)
	if err != nil {
		return "", fmt.Errorf("TranslateText: %v", err)
	}

	translationResult := ""
	// Display the translation for each input text provided
	for _, translation := range resp.GetTranslations() {
		if len(translationResult) > 0 {
			translationResult += "\n"
		}
		translationResult += translation.GetTranslatedText()
	}

	return translationResult, nil
}
