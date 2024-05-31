package count_test

import (
	"os"
	"testing"

	"okcoding.com/grdnr/internal/count"
)

func TestCountWords(t *testing.T) {
	tests := []struct {
		fileName      string
		expectedCount int
	}{
		{
			fileName:      "micro-test.txt",
			expectedCount: 8,
		},
		{
			fileName:      "small-test.txt",
			expectedCount: 916,
		},
	}

	path, err := os.Getwd()
	if err != nil {
		t.Errorf("got error getting current working directory %e", err)
	}

	for _, test := range tests {
		count, err := count.CountWords(path, test.fileName)
		if err != nil {
			t.Errorf("expected no error counting words in %s, got %e", test.fileName, err)
		}

		if count != test.expectedCount {
			t.Errorf("expected %d, got %d", test.expectedCount, count)
		}
	}

}
