package regular

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var founds = []string{}

func WalkAndFindByAlgoAndWord(algo AlgoFunc, word []byte) filepath.WalkFunc {
	f := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("fail on the path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			if IsForbiddenFileExtension(path) {
				return nil
			}
			fileContent, err := os.ReadFile(path)
			// fmt.Println("Reading file: ", path)
			if err != nil {
				return err
			}
			if res := algo(fileContent, word); len(res) != 0 {
				for _, r := range res {
					founds = append(founds, fmt.Sprintf("%v:%v", path, r))
				}
			}
		}
		return nil
	}
	return f
}

func BenchmarkMorisPrattWindowWord(b *testing.B) {
	founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, []byte("window")))
	if len(founds) != 11598 {
		b.Fatal("result did not match with expeceted", len(founds), 11598)
	}
}

func BenchmarkMorisPrattMainWord(b *testing.B) {
	founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, []byte("main")))
	if len(founds) != 19716 {
		b.Fatal("result did not match with expeceted", len(founds), 19716)
	}
}

func BenchmarkMorisPrattFunction(b *testing.B) {
	founds = []string{}
	filepath.Walk(DIR, WalkAndFindByAlgoAndWord(moris_pratt, []byte("function")))
	if len(founds) != 32619 {
		b.Fatal("result did not match with expeceted", len(founds), 32619)
	}
}
