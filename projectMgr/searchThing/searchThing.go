package SearchThing

var BASEDIR string = "/mnt/g/baza_mgr/jacek"

type (
	SearchFileInterface interface {
		FindRegularFile(string) []string
		FindFileByFileExtension(string) []string
		FindFileByRegex(string) []string
		FindFileInZips(string) []string
		FindFileWithSpecialChars(string) []string
	}

	SearchDirInterface interface {
		FindRegularDir(string) []string
		FindDirByFileExtension(string) []string
		FindDirByRegex(string) []string
		FindDirInZips(string) []string
		FindDirWithSpecialChars(string) []string
	}

	SearchTextInterface interface {
		FindRegularText(string) []string
		FindTextByFileExtension(string) []string
		FindTextByRegex(string) []string
		FindTextInZips(string) []string
		FindTextWithSpecialChars(string) []string
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
