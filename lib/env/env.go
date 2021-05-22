package env

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/qyqx233/gtool/lib/files"
)

var NotFound = errors.New("no found")

func GetProjDir() (projDir string, err error) {
	var dir string
	var last string
	var isExist bool
	dir, err = os.Getwd()
	if err != nil {
		return
	}
	for {
		isExist, err = files.Exists(filepath.Join(dir, "go.mod"))
		if err != nil {
			return
		}
		if isExist {
			projDir = dir
			return
		}
		last = dir
		dir = filepath.Dir(dir)
		if last == dir {
			break
		}
	}
	err = NotFound
	return
}
