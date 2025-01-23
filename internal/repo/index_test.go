package repo

import (
	"fmt"
	"os"
	"testing"
)

func TestReadIndexFile(t *testing.T) {
	// arrange
	tPaths := []string{"testcontent.txt", "testcontent2.txt"}
	expected := []IndexEntry{
		{Mode: "100644", Hash: "d7ef1105d426cd3b87a3cf315c763848fd8c7c14", Path: "testcontent.txt"},
		{Mode: "100644", Hash: "2f3c6b82e94acbefbdcc4ac1d00fcfb416892355", Path: "testcontent2.txt"},
	}
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

	// assert
	entries, err := readIndexFile()
	if err != nil {
		t.Fatalf("read index file fail: %v", err)
	}

	for i, entry := range entries {
		if entry != expected[i] {
			t.Fatalf("expected %v but got %v", expected[i], entry)
		}
	}

	// cleanup
	if err := os.RemoveAll(".ggit"); err != nil {
		t.Fatalf("remove .ggit fail: %v", err)
	}

	for _, tPath := range tPaths {
		if err := os.Remove(tPath); err != nil {
			t.Fatalf("cleanup %s fail: %v", tPath, err)
		}
	}

}

func TestTransferToIndexEntry(t *testing.T) {
	// arrange
	testcases := []string{
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent.txt",
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent2.txt",
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent3.txt",
	}
	expected := []IndexEntry{
		{Mode: "100644", Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent.txt"},
		{Mode: "100644", Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent2.txt"},
		{Mode: "100644", Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent3.txt"},
	}

	// action
	entries := transferToIndexEntry(testcases)

	// assert
	for i, entry := range entries {
		if entry != expected[i] {
			t.Fatalf("expected %v but got %v", expected[i], entry)
		}
	}
}
