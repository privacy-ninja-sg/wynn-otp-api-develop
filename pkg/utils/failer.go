package utils

import (
	"github.com/sirupsen/logrus"
	"time"
)

func FailOnErr(err error) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"unix_nano": time.Now().UnixNano(),
		}).Fatal(err)
	}
}

func FailOnErrMsg(err error, msg string) {
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"unix_nano": time.Now().UnixNano(),
			"msg":       msg,
		}).Fatal(err)
	}
}
