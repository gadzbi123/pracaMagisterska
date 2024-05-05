package SearchThing

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

type ProgramFind struct {
	SearchFileInterface
	SearchDirInterface
	name string
}

func NewFind() *ProgramFind {
	return &ProgramFind{name: programExecutableList[Find]}
}

func (p *ProgramFind) FindRegularFile(fileName string) ([]string, error) {
	var out []byte
	var err error
	out, err = exec.Command(p.name, BASEDIR, "-name", fileName).CombinedOutput()
	output := string(out)
	if err != nil {
		log.Println("Error on FindRegularFile", output, err)
		return []string{}, err
	}
	if len(out) == 0 {
		return []string{}, nil
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines, nil
}

func (p *ProgramFind) FindFileByRegex() ([]string, error) {
	return []string{}, errors.New("find can't do regex search in golang")
	/* var out []byte
	var err any
	cmd := exec.Command(programExecutableList[Find], BASEDIR, "-regex", "'.*\\.doc.*'")
	args := cmd.Args
	log.Println(args)
	out, err = cmd.CombinedOutput()
	output := string(out)
	if err != nil {
		log.Println("Error on FindFileByRegex", output, err)
	}
	if len(out) == 0 {
		return []string{}, nil
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines, nil */
}

func FindFileInZips(string) ([]string, error) {
	return []string{}, errors.New("find can't search inside zip")
}
func FindFileWithSpecialChars(string) ([]string, error) {
	return nil, nil
}

func FindFileByPermission(string) ([]string, error) {
	return nil, nil
}
