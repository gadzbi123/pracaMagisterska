package SearchThing

import (
	"testing"
)

func TestFileRegularFind(t *testing.T) {
	fileTester := NewFind()
	result, err := fileTester.FindRegularFile("body-idioms.htm")
	if err != nil {
		t.Error("Failed after regular file execution")
	}
	length := len(result)
	if len(result) == 0 {
		t.Error("The result length was zero")
	}
	t.Log(length)
	t.Log(result)
}

func TestFileNotExisting(t *testing.T) {
	fileTester := NewFind()
	result, err := fileTester.FindRegularFile("Czasy*")
	if err != nil {
		t.Error("Execution of regular file failed")
	}
	length := len(result)
	if len(result) != 0 {
		t.Error("Result length should be zero")
	}
	t.Log(length)
	t.Log(result)
}

// REGEX SUCKS ON find
func TestFileRegexWithFind(t *testing.T) {
	fileTester := NewFind()
	result, err := fileTester.FindFileByRegex()
	if err != nil {
		t.Error("Error during find file by regex")
	}
	length := len(result)
	if length == 0 {
		t.Error("Result length should not be zero")
	}
	t.Log(length)
	t.Log(result)
}

// run tests
// go test -test.v ./*.go
