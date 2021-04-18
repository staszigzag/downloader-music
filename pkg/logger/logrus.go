package logger

import "github.com/sirupsen/logrus"

type Logrus struct {
	Logrus *logrus.Logger
}

func NewLogrus(isDebug bool) *Logrus {
	l := &Logrus{logrus.New()}

	if isDebug {
		// to enable debug
		l.Logrus.Level = logrus.DebugLevel
	}
	// l.Logrus.Formatter = &Logrus.JSONFormatter{}
	return l
}

func (l Logrus) Debug(args ...interface{}) {
	l.Logrus.Debugln(args...)
}

func (l Logrus) Info(args ...interface{}) {
	l.Logrus.Infoln(args...)
}

func (l Logrus) Warn(args ...interface{}) {
	l.Logrus.Warnln(args...)
}

func (l Logrus) Error(args ...interface{}) {
	l.Logrus.Errorln(args...)
}

func (l Logrus) Fatal(args ...interface{}) {
	l.Logrus.Fatalln(args...)
}

func Fatal(args ...interface{}) {
	logrus.Fatalln(args...)
}

func Error(args ...interface{}) {
	logrus.Errorln(args...)
}

func Info(args ...interface{}) {
	logrus.Info(args...)
}
