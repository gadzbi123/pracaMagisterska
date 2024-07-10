package SearchThing

import (
	"errors"
	"log"
	"os/exec"
	"strings"
)

type ProgramFzf struct {
	SearchFileInterface
	SearchDirInterface
	name string
}

func NewFzf() *ProgramFzf {
	n, err := exec.LookPath(programExecutableList[Find])
	if err != nil {
		log.Fatalf("Program %q not found\n", programExecutableList[Find])
	}
	return &ProgramFind{name: n}
}

func (p *ProgramFzf) FindRegularFile(fileName string) ([]string, error) {
	//fzf -1 -q d_test
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

func (p *ProgramFzf) FindFileByRegex() ([]string, error) {
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

func (p *ProgramFzf) FindFileInZips(string) ([]string, error) {
	return []string{}, errors.New("find can't search inside zip")
}

// Double quotes "" are not treated well by linux. They are often converted
// to ' and ' which can be problematic
func (p *ProgramFzf) FindFileWithSpecialChars() ([]string, error) {
	var out []byte
	var err error
	out, err = exec.Command(p.name, BASEDIR, "-name", "*\\'*").CombinedOutput()
	output := string(out)
	if err != nil {
		log.Println("Error on FindSpecialCharsFile", output, err)
		return []string{}, err
	}
	if len(out) == 0 {
		return []string{}, nil
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines, nil
}

func (p *ProgramFzf) FindFileByPermission(perm string) ([]string, error) {
	var out []byte
	var err error
	out, err = exec.Command(p.name, BASEDIR, "-perm", perm).CombinedOutput()
	output := string(out)
	if err != nil {
		log.Println("Error on FindFileByPermission", output, err)
		return []string{}, err
	}
	if len(out) == 0 {
		return []string{}, nil
	}
	//trim last newline
	output_lines := strings.Split(output[:len(output)-1], "\n")
	return output_lines, nil
}
