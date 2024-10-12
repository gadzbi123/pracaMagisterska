package regular

import "path/filepath"

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
