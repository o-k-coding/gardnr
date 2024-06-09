package objectstorage_test

import (
	"context"
	"os"
	"testing"
	"time"

	"okcoding.com/gardnr/internal/testutil"
)

func Test_Cloudflare_CheckFileExists(t *testing.T) {
	storage := testutil.NewTestStorage(t)
	t.Run("check file exists", func(t *testing.T) {
		// this assumes that a file "test.txt" exists in the bucket "test"
		exists, err := storage.CheckFileExists(context.Background(), "test.txt")
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Fatal("text.txt does not exist!")
		}
	})

	t.Run("check file does not exist", func(t *testing.T) {
		exists, err := storage.CheckFileExists(context.Background(), "test-2.txt")
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Fatal("text-2.txt does exist!")
		}
	})
}

func Test_Cloudflare_PutFile(t *testing.T) {
	storage := testutil.NewTestStorage(t)

	t.Run("put a file", func(t *testing.T) {
		// this assumes that a file "test.txt" exists in the bucket "test"
		file, err := os.Open("test.txt")
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()

		err = storage.UploadFile(context.Background(), "test-2.txt", "text/plain", file, time.Minute*1)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			err = storage.DeleteFile(context.Background(), "test-2.txt")
			if err != nil {
				t.Fatal(err)
			}
		})
	})
}
