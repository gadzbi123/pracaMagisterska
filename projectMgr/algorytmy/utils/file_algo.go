package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

const (
	PERF_DIR = "/run/media/gadzbi/GryIFilmy/baza_mgr_large"

	// max_buffer_size = 11 * MB
)

// var global_buffer []byte = make([]byte, max_buffer_size)

type AlgoFunc func([]byte, []byte) []int
type AlgoStruct interface {
	Find([]byte, []byte) []int
}

func IsForbiddenOrArchiveFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf", ".tar.gz", ".rar", ".zip", ".tgz", ".tar", ".gz", ".doc", ".docx":
		return true
	default:
		return false
	}
}
func IsUnsupportedFileExtension(path string) bool {
	unsupportedExt := []string{".ps.gz"}
	for _, ext := range unsupportedExt {
		if strings.HasSuffix(path, ext) {
			return true
		}
	}
	return false
}
func IsForbiddenFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf":
		return true
	default:
		return false
	}
}
func IsGzipExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".tgz", ".gz", ".tar.gz":
		return true
	default:
		return false
	}

}
func isEmptyFile(n int, err error) bool {
	return n == 0 && errors.Is(err, io.EOF)
}

func WalkAndFindByAlgo(algo AlgoStruct, founds *[]string, word []byte) filepath.WalkFunc {
	f := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			slog.Debug("fail on the path", "path", path, "err", err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if IsForbiddenOrArchiveFileExtension(path) {
			return nil
		}
		fd, err := os.Open(path)
		if err != nil {
			slog.Debug("fail to open file", "path", path, "err", err)
			return err
		}
		defer fd.Close()
		scanner := bufio.NewScanner(fd)
		// we don't care about new lines here :)
		for scanner.Scan() {
			if res := algo.Find(scanner.Bytes(), word); len(res) != 0 {
				//think about printing later
				for range res {
					*founds = append(*founds, fmt.Sprintf("%v:%v", path))
				}
			}
		}

		return nil
	}
	return f
}

func PrintFounds(founds []string) {
	base := ""
	for _, f := range founds {
		splitted := strings.Split(f, ":")
		if base == splitted[0] {
			fmt.Printf(":%v", splitted[1])
		} else {
			base = splitted[0]
			fmt.Printf("\n%v", f)
		}
	}
	// fmt.Println()
}
