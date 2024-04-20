package searchThing

type SearchFileInterface interface {
	FindRegularFile(string) []string
	FindFileByFileExtension(string) []string
	FindFileByRegex() []string
	FindFileInZips() []string
	FindFileWithSpecialChars() []string
}

type SearchDirInterface interface {
	FindRegularDir(string) []string
	FindDirByFileExtension(string) []string
	FindDirByRegex() []string
	FindDirInZips() []string
	FindDirWithSpecialChars() []string
}

type SearchTextInterface interface {
	FindRegularText(string) []string
	FindTextByFileExtension(string) []string
	FindTextByRegex() []string
	FindTextInZips() []string
	FindTextWithSpecialChars() []string
}

type SearchByPiping interface {
	todo()
}

type ProgramName int

const (
	Grep ProgramName = iota
	Ripgrep
)

var programExecutableList []string = []string{"grep", "ripgrep"}

type Program struct {
	SearchFileInterface
	SearchDirInterface
	SearchTextInterface
	name ProgramName
}

type SearchFile struct {
	program Program
}

func execCmd(p ProgramName, args ...string) {

}
func (p *Program) FindRegularFile(string) []string {
	switch p.name {
	case Grep:
		execCmd(programExecutableList[Grep])
	case Ripgrep:
	}
	return nil
}

func New(programName ProgramName) *SearchFile {
	sf := &SearchFile{program: Program{name: programName}}
	return sf
}
func (sf *SearchFile) FindRegularFile(s string) {

}
