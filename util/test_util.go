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
func SetConfig() {
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "Sirius") {
		wd = filepath.Dir(wd)
	}
	os.Chdir(wd)
	config.InitConfig("conf/")
}

// function for unit test file get unix time
func GetTimeNowString() string {
	return fmt.Sprintf("\ttime:%s", strconv.FormatInt(time.Now().UnixNano(), 10))
}
