package repo

import (
	"os"
	"testing"
)

func TestInitRepository(t *testing.T) {
	err := InitRepository()
	if err != nil {
		t.Fatalf("Failed to initialize repository: %v", err)
	}

	// Check if .ggit directory exists
	if _, err := os.Stat(".ggit"); os.IsNotExist(err) {
		t.Fatalf(".ggit directory not created")
	}

	// Clean up after test
	err = os.RemoveAll(".ggit")
	if err != nil {
		t.Fatalf("Failed to clean up .ggit directory: %v", err)
	}
}
