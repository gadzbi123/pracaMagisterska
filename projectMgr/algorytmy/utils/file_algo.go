package utils

import (
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

	max_buffer_size = 11 * MB
)

var global_buffer []byte = make([]byte, max_buffer_size)

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
func IsForbiddenFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf":
		return true
	default:
		return false
	}
}

func isEmptyFile(n int, err error) bool {
	return n == 0 && errors.Is(err, io.EOF)
}

func WalkAndFindCmd(algo AlgoStruct, founds *[]string, word []byte, prealloc_buffer *[]byte) filepath.WalkFunc {
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
		if prealloc_buffer != nil {
			pb := *prealloc_buffer
			fd, err := os.Open(path)
			if err != nil {
				slog.Debug("fail to open file", "path", path, "err", err)
				return err
			}
			defer fd.Close()
			n, err := fd.Read(pb)
			if err != nil {
				if isEmptyFile(n, err) {
					return nil
				}
				slog.Debug("fail to read file", "path", path, "err", err)
				return err
			}
			buff := pb[:n]
			if res := algo.Find(buff, word); len(res) != 0 {
				for _, r := range res {
					*founds = append(*founds, fmt.Sprintf("%v:%v", path, r))
				}
			}
			if n, err := fd.Read(pb); !isEmptyFile(n, err) {
				slog.Debug("file was not read fully", "path", path, "err", err, "remaining", n)
				return err
			}
			return nil
		} else {
			buff, err := os.ReadFile(path)
			if err != nil {
				slog.Debug("failed to read file", "error", err)
				return err
			}
			if res := algo.Find(buff, word); len(res) != 0 {
				for _, r := range res {
					*founds = append(*founds, fmt.Sprintf("%v:%v", path, r))
				}
			}
			return nil
		}
	}
	return f
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
		n, err := fd.Read(global_buffer)
		if err != nil {
			if isEmptyFile(n, err) {
				return nil
			}
			slog.Debug("fail to read file", "path", path, "err", err)
			return err
		}
		buff := global_buffer[:n]
		if res := algo.Find(buff, word); len(res) != 0 {
			for _, r := range res {
				*founds = append(*founds, fmt.Sprintf("%v:%v", path, r))
			}
		}
		if n, err := fd.Read(global_buffer); !isEmptyFile(n, err) {
			slog.Debug("file was not read fully", "path", path, "err", err, "remaining", n)
			return err
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
}
