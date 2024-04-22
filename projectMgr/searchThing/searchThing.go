package SearchThing

import (
	"log"
	"os/exec"
	"strings"
)

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
	Find
)

var baseDir string = "/mnt/g/baza_mgr/jacek"

var programExecutableList []string = []string{"grep", "ripgrep", "find"}

type Program struct {
	SearchFileInterface
	SearchDirInterface
	SearchTextInterface
	name ProgramName
}

type SearchFile struct {
	program Program
}

func New(programName ProgramName) *SearchFile {
	sf := &SearchFile{program: Program{name: programName}}
	return sf
}

func (p *Program) FindRegularFile(fileName string) []string {
	var out []byte
	var err any
	switch p.name {
	case Find:
		out, err = exec.Command(programExecutableList[Find], baseDir, "-name", fileName).CombinedOutput()
	case Grep:
		err = "Grep cannot find files"
	}
	output := string(out)
	if err != nil {
		log.Fatalln("Error on FindRegularFile", output, err)
	}
	if len(out) == 0 {
		return []string{}
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines
}

func (p *Program) FindFileByRegex() []string {
	var out []byte
	var err any
	switch p.name {
	case Find:
		cmd := exec.Command(programExecutableList[Find], baseDir, "-regex", "\".*\\.doc.*\"")
		args := cmd.Args
		log.Println(args)
		out, err = cmd.CombinedOutput()
	case Grep:
		err = "Grep cannot find files"
	}
	output := string(out)
	if err != nil {
		log.Fatalln("Error on FindFileByRegex", output, err)
	}
	log.Println("output", out)
	if len(out) == 0 {
		return []string{}
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines
}
