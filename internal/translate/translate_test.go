package translate_test

import (
	"testing"

	"okcoding.com/gardnr/internal/translate"
)

func TestTranslateText(t *testing.T) {
	result, err := translate.TranslateText("en-us", "es", "hello, world!")

	if err != nil {
		t.Fatalf("Expect nil error for TranslateText, got %e", err)
	}

	expect := "¡Hola Mundo!"

	if result != expect {
		t.Fatalf("Expected %s, got %s", expect, result)
	}
}
