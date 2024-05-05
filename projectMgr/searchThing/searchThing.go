package SearchThing

var BASEDIR string = "/mnt/g/baza_mgr/jacek"

type (
	SearchFileInterface interface {
		FindRegularFile(string) ([]string, error)
		FindFileByFileExtension(string) ([]string, error)
		FindFileByRegex(string) ([]string, error)
		FindFileInZips(string) ([]string, error)
		FindFileWithSpecialChars(string) ([]string, error)
	}

	SearchDirInterface interface {
		FindRegularDir(string) ([]string, error)
		FindDirByFileExtension(string) ([]string, error)
		FindDirByRegex(string) ([]string, error)
		FindDirInZips(string) ([]string, error)
		FindDirWithSpecialChars(string) ([]string, error)
	}

	SearchTextInterface interface {
		FindRegularText(string) ([]string, error)
		FindTextByFileExtension(string) ([]string, error)
		FindTextByRegex(string) ([]string, error)
		FindTextInZips(string) ([]string, error)
		FindTextWithSpecialChars(string) ([]string, error)
	}

	SearchByPiping interface {
		todo()
	}
)
type ProgramName int

const (
	Grep ProgramName = iota
	Ripgrep
	Find
)

var programExecutableList []string = []string{"grep", "ripgrep", "find"}
