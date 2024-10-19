package regular

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

const DIR = "/run/media/gadzbi/GryIFilmy/baza_mgr_large"

type AlgoFunc = func([]byte, []byte) []int

func IsForbiddenFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf", ".tar.gz", ".rar", ".zip", ".tgz", ".tar", ".gz":
		return true
	default:
		return false
	}
}

var totalBytes = 0
var totalFiles = 0

func WalkAndFindByAlgoAndWord(algo AlgoFunc, founds *[]string, word []byte) filepath.WalkFunc {
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
			totalBytes += len(fileContent)
			totalFiles++
			if res := algo(fileContent, word); len(res) != 0 {
				for _, r := range res {
					*founds = append(*founds, fmt.Sprintf("%v:%v", path, r))
				}
			}
		}
		return nil
	}
	return f
}
