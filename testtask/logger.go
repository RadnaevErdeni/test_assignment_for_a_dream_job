package testtask

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}
