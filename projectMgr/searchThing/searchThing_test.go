package SearchThing

import (
	"testing"
)

func TestFileRegularFind(t *testing.T) {
	fileTester := New(Find)

	result := fileTester.program.FindRegularFile("body-idioms.htm")
	length := len(result)
	if len(result) == 0 {
		t.Error("Failed on finding regular file")
		t.FailNow()
	}
	t.Log(length)
	t.Log(result)
}

func TestFileNotExisting(t *testing.T) {
	fileTester := New(Find)
	result := fileTester.program.FindRegularFile("Czasy*")
	length := len(result)
	if len(result) != 0 {
		t.Error("Result length should be zero")
		t.FailNow()
	}
	t.Log(length)
	t.Log(result)
}

//REGEX SUCKS ON find
// func TestFileRegexWithFind(t *testing.T) {
// 	fileTester := New(Find)
// 	result := fileTester.program.FindFileByRegex()
// 	length := len(result)
// 	if length == 0 {
// 		t.Error("Result length should not be zero")
// 		t.FailNow()
// 	}
// 	t.Log(length)
// 	t.Log(result)
// }

// run tests
// go test ./*.go
