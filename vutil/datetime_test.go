package vutil

import (
	"time"
	"testing"
)

func TestStrToDate(t *testing.T) {
	datestr := "20180724"
	timestr := "20180724125959"
	date,_ := time.Parse("2018-07-24","2018-07-24")
	time,_:= time.Parse("2018-07-24 01:01:01","2018-07-24 12:59:59")
	dateresult,err := StrToDate(datestr)
	timeresult,err := StrToDate(timestr)
	if err!=nil || dateresult!=date || timeresult != time{
		t.Fail()
	}
}
func TestFormatDate(t *testing.T) {
	date := time.Now()
	datestr:= FormatDate(date,YYYYMMDD)
	timestr:= FormatDate(date,YYYYMMDDHHMMSS)
	_,err:=StrToDate(datestr)
	if err!=nil{
		t.Fail()
	}
	timeresult,err:=StrToDate(timestr)
	if  err != nil && timeresult != date{
		t.Fail()
	}
}