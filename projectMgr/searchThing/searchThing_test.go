package searchThing

import "testing"

func TestFile(t *testing.T) {
	fileTester := SearchFile{}

	result := fileTester.FindRegularFile("xd")
	if len(result) == 0 {
		t.Errorf("Failed on finding regular file")
	}
	t.Log(result)
}

// run tests
// go test ./...
