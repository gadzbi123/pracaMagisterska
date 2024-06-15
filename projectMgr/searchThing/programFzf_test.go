package SearchThing

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileRegularFzf(t *testing.T) {
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

func TestFileNotExistingFzf(t *testing.T) {
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
/* func TestFileRegexWithFind(t *testing.T) {
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
} */

func TestFileWithSpecialCharsInFzf(t *testing.T) {
	fileTester := NewFind()
	result, err := fileTester.FindFileWithSpecialChars()
	if err != nil {
		t.Error("Error during find", err)
	}
	length := len(result)
	if length == 0 {
		t.Error("Result length should not be zero")
	}
	t.Log(length)
	t.Log(result)
}

func TestFileWithSpecifiedPermissionFzf(t *testing.T) {
	fileName := "test-perm.txt"
	filePath := filepath.Join(BASEDIR, fileName)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0555)
	if err != nil {
		t.Error("Couldn't create file ", err)
	}
	file.Close()
	fileTester := NewFind()
	result, err := fileTester.FindFileByPermission("0555")
	if err != nil {
		t.Error("Error during find", err)
	}
	length := len(result)
	if length == 0 {
		t.Error("Result length should not be zero")
	}
	os.Remove(filePath)
	t.Log("Len:", length)
	t.Log("Result:", result)

}

// run tests
// go test -test.v ./*.go
