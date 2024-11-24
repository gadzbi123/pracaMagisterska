package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/regular"
	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

func isInputFromPipe() bool {
	fileInfo, _ := os.Stdin.Stat()
	return fileInfo.Mode()&os.ModeCharDevice == 0
}

func isValidDir(dir string) bool {
	_, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("Not a dir: %v\n", err)
		return false
	}
	return true
}

var ErrInvalidFlag = errors.New("flag provided as substring or file argument")
var ErrNotEnoughArgs = errors.New("expected at least 2 arguments")
var ErrDirWalkFailed = errors.New("walk dir failed")

func isPreallocUsed(static_buffer *bool, buffer_size *string) bool {
	return *static_buffer || *buffer_size != ""
}

// const Usage = `example from gpt`

// Charakterystyka plików (wielkość, ilość)
// Zatrzymanie czasu podczas ładowania z dysku
// Dlaczego język go
// Czy język jest kompilowany do pliku wykonwalnego
// Wykonanie rozruchu po pełnym scachowaniu plików

func run(w io.Writer) error {
	// panic("test real output of the exec")
	flagSet := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	debug := flagSet.Bool("debug", false, "")
	flagSet.BoolVar(debug, "d", *debug, "")

	buffer_size := flagSet.String("buffer", "", "")
	flagSet.StringVar(buffer_size, "b", *buffer_size, "")

	static_buffer := flagSet.Bool("static-buffer", false, "")
	flagSet.BoolVar(static_buffer, "s", *static_buffer, "")

	if err := flagSet.Parse(os.Args[1:]); err != nil {
		return err
	}

	flagSet.Usage = func() {
	}

	lvl := slog.LevelInfo
	if *debug {
		lvl = slog.LevelDebug
	}
	// panic("MAKE SLOG NOT SHOW UP IN TESTS BUT SHOW UP IN MAIN")
	logger := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{Level: lvl}))
	slog.SetDefault(logger)
	slog.Debug("Starting debug logging")

	var prealloc_buffer *[]byte = nil
	if isPreallocUsed(static_buffer, buffer_size) {
		if *buffer_size == "" {
			buff := make([]byte, 512*utils.MB)
			prealloc_buffer = &buff
		} else {
			bs, err := utils.ParseBufferSize(*buffer_size)
			if err != nil {
				slog.Error("invalid buffer size", "buffer", *buffer_size, "error", err)
				flagSet.Usage()
				return err
			}
			buff := make([]byte, bs)
			prealloc_buffer = &buff
		}
	}

	for _, arg := range flagSet.Args() {
		if len(arg) > 0 && arg[0] == '-' {
			slog.Error(ErrInvalidFlag.Error(), "arg", arg)
			flagSet.Usage()
			return ErrInvalidFlag
		}
	}
	var founds = []string{}
	var algo = &regular.BoyerMoore{}
	var args = flagSet.Args()
	switch {
	case len(args) < 2:
		slog.Error("expected at least 2 arguments")
		flagSet.Usage()
		return ErrNotEnoughArgs
	default:
		substring := args[0]
		dir := args[1]
		err := filepath.Walk(dir, utils.WalkAndFindCmd(algo, &founds, []byte(substring), prealloc_buffer))
		if err != nil {
			slog.Warn(ErrDirWalkFailed.Error(), "err", err)
			return ErrDirWalkFailed
		}
	}
	// utils.PrintFounds(founds)
	return nil
}

func main() {
	if err := run(os.Stdout); err != nil {
		os.Exit(1)
	}
}
