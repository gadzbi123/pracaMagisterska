package regular

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

var founds = []string{}

const DIR = "/run/media/gadzbi/GryIFilmy/baza_mgr_large"

func IsForbiddenFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf", ".tar.gz", ".rar", ".zip", ".tgz", ".tar", ".gz":
		return true
	default:
		return false
	}
}

func BenchmarkMorisPrattSimple(b *testing.B) {
	filepath.Walk(DIR, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if !info.IsDir() {
			if IsForbiddenFileExtension(path) {
				return nil
			}
			fileContent, err := os.ReadFile(path)
			fmt.Println("Reading file: ", path)
			if err != nil {
				return err
			}
			if res := moris_pratt(fileContent, []byte("window")); len(res) != 0 {
				for _, r := range res {
					founds = append(founds, string(fileContent[r:r+4]))
				}
			}
		}
		return nil
	})
	fmt.Println("Result:", founds)
}
