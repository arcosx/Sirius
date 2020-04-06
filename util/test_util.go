package util

import (
	"fmt"
	"github.com/arcosx/Sirius/config"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// function for unit test file get right config path
// not use it in production
func SetConfig() {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "Sirius") {
		wd = filepath.Dir(wd)
	}
	os.Chdir(wd)
	config.InitConfig("conf/")
}

func GetTimeNowString() string {
	return fmt.Sprintf("\ttime:%s", strconv.FormatInt(time.Now().UnixNano(), 10))
}
