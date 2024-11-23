package utils

import (
	"errors"
	"strconv"
	"strings"
)

type ByteSize int64

const (
	B  ByteSize = 1
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
)

// Size is made based on size of the biggest file in the dataset

func safeMultiply(a, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, nil
	}

	// Check for overflow
	result := a * b
	if a != result/b {
		return 0, ErrOverflow
	}

	max_prealloc_buffer := 2048 * GB
	if result > int(max_prealloc_buffer) {
		return 0, ErrOverflow
	}

	return result, nil
}
func parseWithExtension(buffer_size string, ext string, ext_value ByteSize) (int, error) {
	numStr := strings.TrimSuffix(buffer_size, ext)
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}
	return safeMultiply(num, int(ext_value))
}
func ParseBufferSize(buffer_size string) (int, error) {
	buffer_size = strings.ToUpper(buffer_size)

	switch {
	case strings.HasSuffix(buffer_size, "GB"):
		return parseWithExtension(buffer_size, "GB", GB)
	case strings.HasSuffix(buffer_size, "MB"):
		return parseWithExtension(buffer_size, "MB", MB)
	case strings.HasSuffix(buffer_size, "KB"):
		return parseWithExtension(buffer_size, "KB", KB)
	case strings.HasSuffix(buffer_size, "B"):
		return parseWithExtension(buffer_size, "B", B)
	default:
		return parseWithExtension(buffer_size, "", B)
	}

}

var ErrOverflow = errors.New("buffer size would overflow integer limits")
