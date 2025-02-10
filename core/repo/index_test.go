package repo

import (
	"fmt"
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	path := "testcontent.txt"
	expectEntry := Entry{
		Mode: 100644,
		Hash: "d7ef1105d426cd3b87a3cf315c763848fd8c7c14",
		Path: path,
	}
	content := fmt.Sprintf("%d %s %s\n", expectEntry.Mode, expectEntry.Hash, expectEntry.Path)
	if err := os.WriteFile("testcontent.txt", []byte(content), 777); err != nil {
		t.Fatalf("fail to create test file: %v", err)
	}

	// action
	entries, err := readFile(path)
	if err != nil {
		t.Fatalf("read index file fail %v", err)
	}
	actualEntry := entries[0]

	// assert
	if actualEntry.Mode != expectEntry.Mode {
		t.Fatalf("expected Mode to be %d, but get %d", expectEntry.Mode, actualEntry.Mode)
	}

	if actualEntry.Hash != expectEntry.Hash {
		t.Fatalf("expected Mode to be %s, but get %s", expectEntry.Hash, actualEntry.Hash)
	}

	if actualEntry.Path != expectEntry.Path {
		t.Fatalf("expected Mode to be %s, but get %s", expectEntry.Path, actualEntry.Path)
	}

	// cleanup
	if err := os.Remove(path); err != nil {
		t.Fatalf("cleanup %s fail: %v", path, err)
	}
}
