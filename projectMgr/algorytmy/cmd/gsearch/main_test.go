package main

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/regular"
	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
	"github.com/gen2brain/go-unarr"
)

/*
func TestIsPreallocUsed(t *testing.T) {
	scenarios := []struct {
		name         string
		staticBuffer bool
		bufferSize   string
		expected     bool
	}{
		{
			name:         "both false or empty",
			staticBuffer: false,
			bufferSize:   "",
			expected:     false,
		},
		{
			name:         "static buffer true",
			staticBuffer: true,
			bufferSize:   "",
			expected:     true,
		},
		{
			name:         "buffer size set",
			staticBuffer: false,
			bufferSize:   "512MB",
			expected:     true,
		},
		{
			name:         "both set",
			staticBuffer: true,
			bufferSize:   "512MB",
			expected:     true,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			staticBuffer := scenario.staticBuffer
			bufferSize := scenario.bufferSize
			result := isPreallocUsed(&staticBuffer, &bufferSize)

			if result != scenario.expected {
				t.Errorf("expected %v but got %v", scenario.expected, result)
			}
		})
	}
}

// TestRun tests the main run function with various scenarios
func TestRun(t *testing.T) {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	// Create a temporary directory for file operations
	tmpDir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test file
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatal(err)
	}

	scenarios := []struct {
		name        string
		args        []string
		expectError error
		skip        bool
	}{
		{
			name:        "no arguments",
			args:        []string{"program"},
			expectError: ErrNotEnoughArgs,
		},
		{
			name:        "invalid flag as argument",
			args:        []string{"program", "pattern", "-invalidflag"},
			expectError: ErrInvalidFlag,
		},
		{
			name:        "valid simple case",
			args:        []string{"program", "test", tmpDir},
			expectError: nil,
		},
		{
			name:        "invalid buffer size",
			args:        []string{"program", "-b", "invalid", "pattern", tmpDir},
			expectError: strconv.ErrSyntax,
		},
		{
			name:        "valid buffer size",
			args:        []string{"program", "-b", "512MB", "pattern", tmpDir},
			expectError: nil,
		},
		{
			name:        "static buffer",
			args:        []string{"program", "-s", "pattern", tmpDir},
			expectError: nil,
		},
		{
			name:        "debug mode",
			args:        []string{"program", "-d", "pattern", tmpDir},
			expectError: nil,
		},
		{
			name:        "debug mode at the end",
			args:        []string{"program", "pattern", tmpDir, "-d"},
			expectError: nil,
			skip:        true,
		},
		{
			name:        "invalid directory",
			args:        []string{"program", "pattern", "/nonexistent/directory"},
			expectError: ErrDirWalkFailed,
		},
	}

	slog.SetDefault(slog.New(slog.NewTextHandler(io.Discard, nil)))

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			// Set up test args
			os.Args = scenario.args

			// Capture logger output
			var buf bytes.Buffer

			// Run optional setup
			if scenario.skip {
				t.Skip("Skipping this test", scenario.name)
			}

			// Run the function
			err := run(&buf)

			// Check error
			if !errors.Is(err, scenario.expectError) {
				t.Errorf("expected error %v, got %v", scenario.expectError, err)
			}
		})
	}
}

// TestBufferPreallocation tests the buffer preallocation logic
func TestBufferPreallocation(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	scenarios := []struct {
		name          string
		args          []string
		expectedSize  int
		expectNilBuff bool
		expectError   error
	}{
		{
			name:          "no buffer",
			args:          []string{"program", "pattern", "."},
			expectNilBuff: true,
			expectError:   nil,
		},
		{
			name:          "static buffer",
			args:          []string{"program", "-s", "pattern", "."},
			expectedSize:  512 * int(utils.MB),
			expectNilBuff: false,
			expectError:   nil,
		},
		{
			name:          "custom buffer size",
			args:          []string{"program", "-b", "1MB", "pattern", "."},
			expectedSize:  1 * int(utils.MB),
			expectNilBuff: false,
			expectError:   nil,
		},
		{
			name:          "invalid buffer size",
			args:          []string{"program", "-b", "invalid", "pattern", "."},
			expectNilBuff: true,
			expectError:   strconv.ErrSyntax,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			os.Args = scenario.args

			// Capture logger output
			var buf bytes.Buffer
			err := run(&buf)

			// Check error
			if !errors.Is(err, scenario.expectError) {
				t.Errorf("expected error %v, got %v", scenario.expectError, err)
			}

			if scenario.expectError != nil {
				if err == nil {
					t.Error("expected error but got none")
				}
			}
		})
	}
}

// TestDebugFlag tests the debug flag functionality
func TestDebugFlag(t *testing.T) {
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	scenarios := []struct {
		name          string
		args          []string
		expectedLevel slog.Level
	}{
		{
			name:          "debug enabled",
			args:          []string{"program", "-d", "pattern", "."},
			expectedLevel: slog.LevelDebug,
		},
		{
			name:          "debug disabled",
			args:          []string{"program", "pattern", "."},
			expectedLevel: slog.LevelInfo,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			os.Args = scenario.args

			// Capture logger output
			var buf bytes.Buffer

			// Run function
			_ = run(&buf)

			// Check log level in output
			// Note: This is a simplified check and might need adjustment
			// based on your actual logging implementation
			if scenario.expectedLevel == slog.LevelDebug {
				if !bytes.Contains(buf.Bytes(), []byte("DEBUG")) {
					t.Error("expected debug level logging but got none")
				}
			}
		})
	}
}
*/

func TestArchiveMain(t *testing.T) {
	t.SkipNow()
	path := LIB_DIR
	dirEnt, err := os.ReadDir(path)
	if err != nil {
		t.Error(err)
	}
	for _, d := range dirEnt {
		di, err := d.Info()
		if err != nil {
			t.Error(err)
		}
		// err = SearchArchive(path + di.Name())
		if err != nil {
			t.Errorf("failed to read archive (%v): %v\n", di.Name(), err)
		}
	}
}

const LIB_DIR = "/run/media/gadzbi/GryIFilmy/baza_mgr/"

func TestExt(t *testing.T) {
	text := "/tmp/baza_mgr_unarr/polskie/polskie/www/Html/Html/doc/DA.doc"
	if filepath.Ext(text) == ".doc" {
		return
	}
	t.Fail()
}

func TestArchiveSummary(t *testing.T) {
	t.SkipNow()
	// debug.SetMemoryLimit(8*2 ^ 30)
	// sudo mount -o remount,size=30G /tmp/
	os.RemoveAll("/tmp/baza_mgr_unarr")
	algo := regular.BoyerMoore{}
	err := filepath.Walk(LIB_DIR, searchInFiles(&algo, []byte("window")))
	if err != nil {
		t.Error(err)
	}

}
func searchInFiles2(algo *regular.BoyerMoore, word []byte) filepath.WalkFunc {
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
				// Create a tar reader
				tarReader := tar.NewReader(gzReader)
			}
		*/
		if utils.IsArchiveFileExtension(path) && utils.IsNotKnownCorruptedFile(path) {
			err = UnzipArchive(path, algo, word)
			return err
		}
		fd, err := os.Open(path)
		if err != nil {
			fmt.Printf("fail to open file", "path", path, "err", err)
			return err
		}
		defer fd.Close()
		scanner := bufio.NewScanner(fd)
		// we don't care about new lines here :)
		line := 1
		for scanner.Scan() {
			if res := algo.Find(scanner.Bytes(), word); len(res) != 0 {
				for range res {
					// fmt.Printf("%v:l%v:c%v\n", path, line, char)
				}
			}
			line++
		}
		if err != nil && err != io.EOF {
			fmt.Printf("failed to scan (%v): %v\n", path, err)
			return err
		}
		return nil
	}
	return f
}
func UnzipArchive2(path string, algo *regular.BoyerMoore, word []byte) error {
	fmt.Printf("reading new archive: %v\n", path)
	a, err := unarr.NewArchive(path)
	if err != nil {
		fmt.Printf("failed to read inner archive (%v): %v\n", path, err)
		return nil
	}
	defer a.Close()
	base_ext := filepath.Base(path)
	temp_path := filepath.Join(os.TempDir(), "baza_mgr_unarr")
	new_path, _ := strings.CutPrefix(path, LIB_DIR)
	new_path, _ = strings.CutSuffix(new_path, filepath.Ext(base_ext))
	if !strings.HasPrefix(new_path, temp_path) {
		new_path = filepath.Join(temp_path, new_path)
	}
	new_path += ".unzipped"
	fmt.Printf("extracting archive to: %v\n", new_path)
	_, err = a.Extract(new_path)
	if err != nil {
		fmt.Printf("failed to extract inner archive (%v): %v\n", path, err)
		return nil
	}

	err = filepath.Walk(new_path, searchInFiles(algo, word))
	if err != nil {
		fmt.Printf("inner walk failed (%v): %v", new_path, err)
		return err
	}
	return nil
}

/*
from MEMORY

func TestArchiveMain(t *testing.T) {
	t.SkipNow()
	debug.SetMemoryLimit(16*2 ^ 30)
	path := "/run/media/gadzbi/GryIFilmy/baza_mgr/"
	dirEnt, err := os.ReadDir(path)
	if err != nil {
		t.Error(err)
	}
	for _, d := range dirEnt {
		di, err := d.Info()
		if err != nil {
			t.Error(err)
		}
		err = SearchArchive(path + di.Name())
		if err != nil {
			t.Errorf("failed to read archive (%v): %v\n", di.Name(), err)
		}
	}
}
func TestArchiveSummary(t *testing.T) {
	// t.SkipNow()
	err := SearchArchive("/run/media/gadzbi/GryIFilmy/baza_mgr/angielski.zip")
	if err != nil {
		t.Error(err)
	}

}
func SearchArchive(path string) error {
	a, err := unarr.NewArchive(path)
	if err != nil {
		return err
	}
	defer a.Close()
	err = SearchArchiveRec(a, "/")
	return err
}
func SearchArchiveRec(a *unarr.Archive, rootDir string) error {
	for {
		err := a.Entry()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("invalid entry (%v): %v\n", a.Name(), err)
				return err
			}
		}
		if a.Size() == 0 {
			continue
		}
		name := a.Name()
		if utils.IsUnsupportedFileExtension(name) {
			fmt.Printf("unsupported file %v\n", name)
			continue
		}
		content, err := a.ReadAll()
		if err != nil {
			fmt.Printf("failed to read all of file data: %v\n", a.Name())
			continue
		}
		// if utils.IsMicrosoftFileExtension(name) {
		// 	zip.NewReader()

		// }
		if utils.IsArchiveFileExtension(name) {
			err = a.EntryFor(name)
			if err != nil {
				fmt.Printf("failed to do entryFor archive (%v): %v\n", name, err)
				continue
			}

			fmt.Printf("reading new archive: %v\n", name)
			a2, err := unarr.NewArchiveFromMemory(content)
			if err != nil {
				fmt.Printf("failed to read inner archive (%v): %v\n", name, err)
				continue
			}
			defer a2.Close()
			// a.Entry()
			err = SearchArchiveRec(a2, rootDir+name+"/")
			// if err != nil {
			// 	fmt.Printf("failed to read inner archive: %v\n", err)
			// }
			continue
		}
		bm := regular.BoyerMoore{}
		res := bm.Find(content, []byte("main"))
		if len(res) != 0 {
			// fmt.Printf("%v:%v\n", rootDir+name, res)
		}
	}

	return nil
}
*/
