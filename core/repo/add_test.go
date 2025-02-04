package repo

import (
	"os"
	"testing"
)

func TestAddFileToIndex(t *testing.T) {
	// arrange
	tPath := "testcontent.txt"
	if err := os.WriteFile(tPath, []byte("test"), 0777); err != nil {
		t.Fatalf("fail to create test file: %v", err)
	}

	// action
	if err := InitRepository(); err != nil {
		t.Fatalf("init .ggit fail: %v", err)
	}

	if err := AddFileToIndex(tPath); err != nil {
		t.Fatalf("adding file fail %v", err)
	}

	// cleanup the temp file
	if err := os.RemoveAll(".ggit"); err != nil {
		t.Fatalf("remove .ggit fail: %v", err)
	}

	if err := os.Remove(tPath); err != nil {
		t.Fatalf("cleanup %s fail: %v", tPath, err)
	}
}
