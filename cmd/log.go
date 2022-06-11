package main

import (
	"fmt"
	"io"
	"os"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
)

func main() {

	logName := fmt.Sprintf("./log/access_log.")
	r, _ := rotatelogs.New(logName + "%Y%m%d")
	mw := io.MultiWriter(os.Stdout, r)
	log.SetOutput(mw)
	log.Info("something ....")
}
