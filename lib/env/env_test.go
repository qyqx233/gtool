package env

import (
	"os"
	"path"
	"testing"
)

func Test_env1(t *testing.T) {
	wd, _ := os.Getwd()
	t.Log(path.Split(wd))
	// t.Log(strings.Split(wd, os.PathSeparator))
	dir, err := GetProjDir()

	t.Log(dir, err)
}
