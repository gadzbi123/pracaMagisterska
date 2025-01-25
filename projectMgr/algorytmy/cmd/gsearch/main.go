package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/regular"
	"github.com/gen2brain/go-unarr"
)

// func isInputFromPipe() bool {
// 	fileInfo, _ := os.Stdin.Stat()
// 	return fileInfo.Mode()&os.ModeCharDevice == 0
// }

var ErrNotEnoughArgs = errors.New("expected at least 2 arguments")
var ErrDirWalkFailed = errors.New("walk dir failed")
var TMP_DIR = "/tmp/baza_mgr_unarr"
var START_ROOT = ""

// const Usage = `example from gpt`

// Dlaczego język go
// Czy język jest kompilowany do pliku wykonwalnego
// Wykonanie rozruchu po pełnym scachowaniu plików

// go run cmd/gsearch/main.go function /run/media/gadzbi/GryIFilmy/baza_mgr_small/papers/jmlr_specjal/jmlr/volume2/crammer01a
// rg -lcU function /run/media/gadzbi/GryIFilmy/baza_mgr_small/papers/jmlr_specjal/jmlr/volume2/crammer01a
// grep -rnv function /run/media/gadzbi/GryIFilmy/baza_mgr_small/papers/jmlr_specjal/jmlr/volume2/crammer01a
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

type SearchData struct {
	root_path string
	curr_path string
	algo      *regular.BoyerMoore
	word      []byte
}

func run() error {
	os.RemoveAll(TMP_DIR)

	switch {
	case len(os.Args) < 2:
		fmt.Printf("gsearch: %v", ErrNotEnoughArgs)
		return ErrNotEnoughArgs
	default:
		substring := os.Args[1]
		START_ROOT = os.Args[2]
		bm := &regular.BoyerMoore{}
		err := filepath.Walk(START_ROOT, searchInFiles(START_ROOT, bm, []byte(substring)))
		if err != nil {
			fmt.Printf("%v: %v\n", ErrDirWalkFailed.Error(), err)
			return ErrDirWalkFailed
		}
	}
	return nil
}

func searchInFiles(path string, algo *regular.BoyerMoore, word []byte) filepath.WalkFunc {
	f := func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		/*
			if utils.IsGzipExtension(path) {
				file, err := os.Open(path)
				if err != nil {
					fmt.Printf("failed to open gzip file (%v): %v\n", path, err)
					return err
				}
				defer file.Close()
				gzReader, err := gzip.NewReader(file)
				if err != nil {
					return err
				}
				defer gzReader.Close()
				// os.Create()
				io.ReadAll(gzReader)
				tarReader := tar.NewReader(gzReader)
			}
		*/
		if IsArchiveFileExtension(path) {
			if IsKnownCorruptedFile(path) {
				// fmt.Printf("skipping corrupted file: %v\n", path)
				return nil
			}
			err = UnzipArchive(path, algo, word)
			return err
		}
		fd, err := os.Open(path)
		if err != nil {
			// fmt.Printf("fail to open file (%v): %v\n", path, err)
			return err
		}
		defer fd.Close()
		scanner := bufio.NewScanner(fd)
		// we don't care about new lines here :)
		line := 1
		for scanner.Scan() {
			if res := algo.Find(scanner.Bytes(), word); len(res) != 0 {
				fmt.Printf("%v:l%v\n", path, line)
				// for _, char := range res {
				// 	fmt.Printf("%v:l%v:c%v\n", path, line, char)
				// }
			}
			line++
		}
		if err := scanner.Err(); err != nil {
			// fmt.Printf("failed to scan (%v): %v\n", path, err)
		}
		return nil
	}
	return f
}
func UnzipArchive(path string, algo *regular.BoyerMoore, word []byte) error {
	// fmt.Printf("reading new archive: %v\n", sd.curr_path)
	a, err := unarr.NewArchive(path)
	if err != nil {
		// fmt.Printf("failed to read inner archive (%v): %v\n", path, err)
		return nil
	}
	defer a.Close()
	new_path, _ := strings.CutPrefix(path, START_ROOT)
	new_path, _ = strings.CutSuffix(new_path, filepath.Ext(path))
	if !strings.HasPrefix(new_path, TMP_DIR) {
		new_path = filepath.Join(TMP_DIR, new_path)
	}
	new_path = new_path + ".unzipped"

	// fmt.Printf("extracting archive to: %v\n", sd.curr_path)
	_, err = a.Extract(new_path)
	if err != nil {
		// fmt.Printf("failed to extract inner archive (%v): %v\n", new_path, err)
		return nil
	}

	err = filepath.Walk(new_path, searchInFiles(new_path, algo, word))
	if err != nil {
		fmt.Printf("inner walk failed (%v): %v", new_path, err)
		return err
	}
	return nil
}

func IsArchiveFileExtension(path string) bool {
	ext := filepath.Ext(path)
	switch ext {
	case ".tar.gz", ".rar", ".zip", ".gz", ".tar", ".tgz", ".docx": //".doc"
		return true
	default:
		return false
	}
}

func IsKnownCorruptedFile(path string) bool {
	knownCorruptedFiles := []string{
		"DataMining.zip",
		"angielski.zip",
		"konferencja.zip",
		"literatura_old.zip",
		"matematyka.zip",
		"en_curriculum6_v1.zip",
		"_vti_cnf/allbmp.zip",
		"/_vti_cnf/allswf.zip",
		"Analysis of Incomplete Multivariate Data.zip",
		"[ Neural Networks and Learning - Ebook ] Introduction to Machine Learning - Michael Burl.zip",
		"Practical Handbook of Genetic Algorithms Complex Coding Systems  Volume III.zip",
		"STL_doc.tar.gz",
		"/stl/",
		"Fizyka Wyklady All.rar",
		"Wirusy_Adam_Blaszczyk.rar",
		"2124C Allfiles.zip",
	}
	for _, f := range knownCorruptedFiles {
		if strings.Contains(path, f) {
			return true
		}
	}
	return false
}
