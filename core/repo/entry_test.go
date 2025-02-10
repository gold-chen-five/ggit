package repo

import "testing"

func TestToEntries(t *testing.T) {
	// arrange
	testcases := []string{
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent.txt",
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent2.txt",
		"100644 3b18e512dba79e4c8300dd08aeb37f8e728b8dad testcontent3.txt",
	}
	expected := []Entry{
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent.txt"},
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent2.txt"},
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent3.txt"},
	}

	// action
	entries, err := ToEntries(testcases)
	if err != nil {
		t.Fatalf("toEntries error: %v", err)
	}

	// assert
	for i, entry := range entries {
		if entry != expected[i] {
			t.Fatalf("expected %v but got %v", expected[i], entry)
		}
	}
}

func TestFindEntry(t *testing.T) {
	// arrange
	entries := []Entry{
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent.txt"},
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent2.txt"},
	}
	entries2 := []Entry{
		{Mode: 100644, Hash: "3b18e512dba79e4c8300dd08aeb37f8e728b8dad", Path: "testcontent.txt"},
	}

	// action
	index, isFound := FindEntry(entries, "testcontent2.txt")
	index2, isFound2 := FindEntry(entries, "testcontent3.txt")
	index3, isFound3 := FindEntry(entries2, "testcontent.txt")

	// assert
	if index != 1 || !isFound {
		t.Fatalf("expected 1 but got %d, expected true but got %t", index, isFound)
	}

	if index2 != 2 || isFound2 {
		t.Fatalf("expected 2 but got %d, expected false but got %t", index2, isFound2)
	}

	if index3 != 0 || !isFound3 {
		t.Fatalf("expected 0 but got %d, expected true but got %t", index3, isFound3)
	}
}
