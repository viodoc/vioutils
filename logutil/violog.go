package logutil

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var Log *logrus.Logger

func NewLogger(infoPath,errorPath string) *logrus.Logger {
	if Log != nil {
		return Log
	}

	//pathMap := lfshook.PathMap{
	//	logrus.InfoLevel:  infoPath,
	//	logrus.ErrorLevel: errorPath,
	//}
	infoWriter,err := rotatelogs.New(
		infoPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(infoPath),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
		)
	if err!=nil{
		Log.Fatalf("init log error: %v",err)
	}
	errWriter,err := rotatelogs.New(
		infoPath+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(errorPath),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(604800)*time.Second),
	)
	if err!=nil{
		Log.Fatalf("init log error: %v",err)
	}
	Log = logrus.New()
	//Log.Hooks.Add(lfshook.NewHook(
	//	pathMap,
	//	&logrus.JSONFormatter{},
	//))
	Log.Hooks.Add(lfshook.NewHook(
		lfshook.WriterMap{
			logrus.InfoLevel:  infoWriter,
			logrus.ErrorLevel: errWriter,
		},
		&logrus.JSONFormatter{},
	))
	return Log
}
