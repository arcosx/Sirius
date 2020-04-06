package config

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_new(t *testing.T) {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "Sirius") {
		wd = filepath.Dir(wd)
	}
	os.Chdir(wd)
	InitConfig("conf/")
}