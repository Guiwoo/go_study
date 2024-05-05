package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewCustomLog() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.TextFormatter{})
	log.SetOutput(os.Stdout)
	return log
}
