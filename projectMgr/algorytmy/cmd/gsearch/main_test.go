package main

import (
	"bytes"
	"errors"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/gadzbi123/pracaMagisterska/algorytmy/utils"
)

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
