package tool

import (
	"os"
	"testing"
)

func TestStoreObject(t *testing.T) {
	// arrange
	tContent := "test content"

	// action
	_, err := StoreObject(Commit, tContent)
	if err != nil {
		t.Fatalf("store object fail: %v", err)
	}

	// cleanup the temp file
	if err := os.RemoveAll(".ggit"); err != nil {
		t.Fatalf("remove .ggit fail: %v", err)
	}
}

func TestCompressObject(t *testing.T) {
	// arrange
	tContent := "test content"

	// action
	buf, err := CompressObject(tContent)
	if err != nil {
		t.Fatalf("compress object fail: %v", err)
	}

	// assert
	if len(buf) == 0 {
		t.Fatalf("compress object fail: empty buffer")
	}
}
