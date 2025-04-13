package main

import (
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

func main() {
	var log = logrus.New()
	//log.SetFormatter(&logrus.JSONFormatter{})
	//log.SetLevel(logrus.DebugLevel)

	log.AddHook(&Hook{})

	log.Infoln(errorCounter)

	log.Debug("Debug")
	log.Error("Error")
	log.Warn("Warn")

	log.Infoln(errorCounter)

	l := log.WithField("request_id", 23123)
	l.Info("hello")
	l.Info("process request")

	l2 := l.WithField("test", "test")
	l2.Infof("Test test")
}

var errorCounter, warningCounter uint64

// Hook for logrus.
type Hook struct {
}

// Fire executes hook.
func (hook *Hook) Fire(entry *logrus.Entry) error {
	switch entry.Level {
	case logrus.PanicLevel:
		atomic.AddUint64(&errorCounter, 1)

	case logrus.FatalLevel:
		atomic.AddUint64(&errorCounter, 1)

	case logrus.ErrorLevel:
		atomic.AddUint64(&errorCounter, 1)

	case logrus.WarnLevel:
		atomic.AddUint64(&warningCounter, 1)
	}
	return nil
}

// Levels returns a slice of logrus levels with witch hook work.
func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
	}
}
