package vutil

import (
	"time"
	"bytes"
	"strings"
	"errors"
)
type DateStyle string

const (
	MM_DD                           = "MM-dd"
	YYYYMM                          = "yyyyMM"
	YYYY_MM                         = "yyyy-MM"
	YYYY_MM_DD                      = "yyyy-MM-dd"
	YYYYMMDD                        = "yyyyMMdd"
	YYYYMMDDHHMMSS                  = "yyyyMMddHHmmss"
	YYYYMMDDHHMM                    = "yyyyMMddHHmm"
	YYYYMMDDHH                      = "yyyyMMddHH"
	YYMMDDHHMM                      = "yyMMddHHmm"
	MM_DD_HH_MM                     = "MM-dd HH:mm"
	MM_DD_HH_MM_SS                  = "MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM                = "yyyy-MM-dd HH:mm"
	YYYY_MM_DD_HH_MM_SS             = "yyyy-MM-dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS         = "yyyy-MM-dd HH:mm:ss.SSS"

	MM_DD_EN                        = "MM/dd"
	YYYY_MM_EN                      = "yyyy/MM"
	YYYY_MM_DD_EN                   = "yyyy/MM/dd"
	MM_DD_HH_MM_EN                  = "MM/dd HH:mm"
	MM_DD_HH_MM_SS_EN               = "MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_EN             = "yyyy/MM/dd HH:mm"
	YYYY_MM_DD_HH_MM_SS_EN          = "yyyy/MM/dd HH:mm:ss"
	YYYY_MM_DD_HH_MM_SS_SSS_EN      = "yyyy/MM/dd HH:mm:ss.SSS"

	MM_DD_CN                        = "MM月dd日"
	YYYY_MM_CN                      = "yyyy年MM月"
	YYYY_MM_DD_CN                   = "yyyy年MM月dd日"
	MM_DD_HH_MM_CN                  = "MM月dd日 HH:mm"
	MM_DD_HH_MM_SS_CN               = "MM月dd日 HH:mm:ss"
	YYYY_MM_DD_HH_MM_CN             = "yyyy年MM月dd日 HH:mm"
	YYYY_MM_DD_HH_MM_SS_CN          = "yyyy年MM月dd日 HH:mm:ss"

	HH_MM                           = "HH:mm"
	HH_MM_SS                        = "HH:mm:ss"
	HH_MM_SS_MS                     = "HH:mm:ss.SSS"
)
func StrToDate(date string) (time.Time,error){
	if len(date)==8{
		y := string([]rune(date)[:4])
		m := string([]rune(date)[4:6])
		d := string([]rune(date)[6:8])
		b := bytes.Buffer{}
		b.WriteString(y)
		b.WriteString("-")
		b.WriteString(m)
		b.WriteString("-")
		b.WriteString(d)
		date = b.String()
		return time.Parse("2006-01-02",date)
	}
	if len(date)==14{
		y := string([]rune(date)[:4])
		m := string([]rune(date)[4:6])
		d := string([]rune(date)[6:8])
		h := string([]rune(date)[8:10])
		mm := string([]rune(date)[10:12])
		s := string([]rune(date)[12:14])
		b := bytes.Buffer{}
		b.WriteString(y)
		b.WriteString("-")
		b.WriteString(m)
		b.WriteString("-")
		b.WriteString(d)
		b.WriteString(" ")
		b.WriteString(h)
		b.WriteString(":")
		b.WriteString(mm)
		b.WriteString(":")
		b.WriteString(s)
		date = b.String()
		return time.Parse("2006-01-02 15:04:05",date)
	}
	return time.Time{},errors.New("date/time format error")
}
//日期转字符串
func FormatDate(date time.Time, dateStyle DateStyle) string {
	layout := string(dateStyle)
	layout = strings.Replace(layout, "yyyy", "2006", 1)
	layout = strings.Replace(layout, "yy", "06", 1)
	layout = strings.Replace(layout, "MM", "01", 1)
	layout = strings.Replace(layout, "dd", "02", 1)
	layout = strings.Replace(layout, "HH", "15", 1)
	layout = strings.Replace(layout, "mm", "04", 1)
	layout = strings.Replace(layout, "ss", "05", 1)
	layout = strings.Replace(layout, "SSS", "000", -1)

	return date.Format(layout)
}