package fs

import (
	"bufio"
	"io/ioutil"
	"os"
)

type Filer interface {
	Read(filepath string) ([]byte, error)
	ReadLineByLine(filepath string, walker func(line string)) error
}

type fs struct{}

func New() fs {
	return fs{}
}

func (f fs) Read(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

func (f fs) ReadLineByLine(filepath string, walker func(line string)) error {
	file, err := os.Open(filepath)
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		walker(line)
	}

	return nil
}
