package env

import (
	"strconv"
	"strings"

	"github.com/go-mservice-bench/lib/fs"
)

type Env struct {
	Data map[string]string
}

func New(fs fs.Filer) (e Env, err error) {
	e.Data = map[string]string{}
	err = fs.ReadLineByLine(".env", e.walker())
	return e, err
}

func (e Env) Get(key string) (string, bool) {
	value, exists := e.Data[key]
	return value, exists
}

func (e Env) GetInt(key string) (int, bool) {
	valueStr, exists := e.Get(key)
	if !exists {
		return 0, false
	}
	valueInt, err := strconv.Atoi(valueStr)
	return valueInt, err == nil
}

func (e *Env) walker() func(line string) {
	return func(line string) {
		if !strings.Contains(line, "=") {
			return
		}
		if len(line) == 0 || line[0:1] == "#" {
			return
		}
		sli := strings.Split(line, "=")
		if len(sli) == 1 {
			return
		}
		key := sli[0]
		value := strings.Replace(sli[1], "\"", "", -1)
		e.Data[key] = value
	}
}
