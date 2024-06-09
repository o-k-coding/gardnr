package create

import (
	"context"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/google/uuid"
	"okcoding.com/gardnr/internal/testutil"
)

func Test_handleImage(t *testing.T) {
	// TODO move this out to a common test helper and get the config values from environment
	storage := testutil.NewTestStorage(t)
	t.Run("handle image that has already been uploaded", func(t *testing.T) {
		text := "![[kylar.png]]"
		transformedText, err := handleImage(context.Background(), "./", text, storage)
		if err != nil {
			t.Fatal(err)
		}
		if transformedText != "<Image src='https://images.okcoding.io/kylar.png' alt='kylar.png' inferSize={true}/>" {
			t.Fatal("handleImage did not transform the text")
		}
	})

	t.Run("handle image that has not already been uploaded", func(t *testing.T) {
		// TODO copy kylar.png into a new image with a random id in the file name
		// TODO clean up the files
		id := uuid.New().String()
		file := "kylar" + id + ".png"
		err := copyFile("./kylar.png", "./"+file)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(func() {
			deleteFile("./" + file)
			storage.DeleteFile(context.Background(), file)
		})
		exists, err := storage.CheckFileExists(context.Background(), file)
		if err != nil {
			t.Fatal(err)
		}
		if exists {
			t.Fatalf("%s already exists!", file)
		}
		text := "![[" + file + "]]"
		transformedText, err := handleImage(context.Background(), "./", text, storage)
		if err != nil {
			t.Fatal(err)
		}
		if transformedText != "<Image src='https://images.okcoding.io/"+file+"' alt='"+file+"' inferSize={true}/>" {
			t.Fatal("handleImage did not transform the text")
		}
		exists, err = storage.CheckFileExists(context.Background(), file)
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Fatalf("%s doesn't exist in storage!", file)
		}
	})
}

func deleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func copyFile(src, dst string) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return err
}
