package main

import (
	"io"
	"testing"
)

func TestWriteTape(t *testing.T) {
	file, clean := createTempFile(t, "tape")
	defer clean()

	tape := &tape{file}

	tape.Write([]byte("abc"))

	file.Seek(0, io.SeekStart)
	newFileContents, _ := io.ReadAll(file)

	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
