package regular

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const DIR = "/run/media/gadzbi/GryIFilmy/baza_mgr_large"

type AlgoFunc func([]byte, []byte) []int
type AlgoStruct interface {
	Find([]byte, []byte) []int
}

func IsForbiddenFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".jpg", ".gif", ".pdf", ".tar.gz", ".rar", ".zip", ".tgz", ".tar", ".gz":
		return true
	default:
		return false
	}
}

type ByteSize int64

const (
	B  ByteSize = 1
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

// Size is made based on size of the biggest file in the dataset
const max_buffer_size = 11 * MB

var global_buffer []byte = make([]byte, max_buffer_size)

func isEmptyFile(n int, err error) bool {
	return n == 0 && errors.Is(err, io.EOF)
}

func WalkAndFindByAlgo(algo AlgoStruct, founds *[]string, word []byte) filepath.WalkFunc {
	f := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("fail on the path %q: %v\n", path, err)
			return err
		}
		if info.IsDir() {
			return nil
		}
		if IsForbiddenFileExtension(path) {
			return nil
		}
		fd, err := os.Open(path)
		if err != nil {
			fmt.Printf("Failed to open file=%v, err=%v\n", path, err)
			return err
		}
		defer fd.Close()
		n, err := fd.Read(global_buffer)
		if err != nil {
			if isEmptyFile(n, err) {
				return nil
			}
			fmt.Printf("Failed to read file=%v, err=%v\n", path, err)
			return err
		}
		buff := global_buffer[:n]
		if res := algo.Find(buff, word); len(res) != 0 {
			for _, r := range res {
				*founds = append(*founds, fmt.Sprintf("%v:%v", path, r))
			}
		}
		if n, err := fd.Read(global_buffer); !isEmptyFile(n, err) {
			fmt.Printf("File was not read fully: err=%v, n=%v\n", err, n)
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
