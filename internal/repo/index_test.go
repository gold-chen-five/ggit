package repo

import (
	"fmt"
	"os"
	"testing"
)

func TestReadIndexFile(t *testing.T) {
	// arrange
	tPaths := []string{"testcontent.txt", "testcontent2.txt"}
	for i, tPath := range tPaths {
		if err := os.WriteFile(tPath, []byte(fmt.Sprintf("test %d", i)), 0777); err != nil {
			t.Fatalf("fail to create test file: %v", err)
		}
	}

	// action
	if err := InitRepository(); err != nil {
		t.Fatalf("init .ggit fail: %v", err)
	}

	for _, tPath := range tPaths {
		if err := AddFileToIndex(tPath); err != nil {
			t.Fatalf("adding file fail %v", err)
		}
	}

	lines, err := readIndexFile()
	if err != nil {
		t.Fatalf("read index file fail: %v", err)
	}

	for _, line := range lines {
		t.Fatalf(line)
	}
	// assert
}
