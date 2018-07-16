package logutil

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger

func NewLogger(infoPath,errorPath,level string) *logrus.Logger {
	if Log != nil {
		return Log
	}
	infoWriter,err := rotatelogs.New(
		infoPath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
		)
	if err!=nil{
		logrus.Fatalf("init log error: %v",err)
	}
	errWriter,err := rotatelogs.New(
		errorPath+".%Y%m%d",
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err!=nil{
		logrus.Fatalf("init log error: %v",err)
	}
	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errWriter,
		},
		&logrus.JSONFormatter{},
	))
	lel,err:=logrus.ParseLevel(level)
	Log.SetLevel(lel)
	return Log
}
