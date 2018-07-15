package dbutils

import (
	"github.com/jmoiron/sqlx"
	"time"
)

func namedGet(stmt *sqlx.NamedStmt, dest interface{}, arg interface{}) error {
	return stmt.Get(dest, arg)
}

func get(queryer sqlx.Queryer, dest interface{}, query string, args ...interface{}) error {
	return sqlx.Get(queryer, dest, query, args...)
}

func namedSelect(stmt *sqlx.NamedStmt, dest interface{}, arg interface{}) {
	err := stmt.Select(dest, arg)
	panicErr(err)
}

func selectx(queryer sqlx.Queryer, dest interface{}, query string, args ...interface{}) {
	err := sqlx.Select(queryer, dest, query, args...)
	panicErr(err)
}

func exec(e sqlx.Execer, query string, args ...interface{}) Result {
	result := sqlx.MustExec(e, query, args...)
	return &vioResult{result}
}

func namedExec(ext sqlx.Ext, query string, arg interface{}) Result {
	result, err := sqlx.NamedExec(ext, query, arg)
	panicErr(err)
	return &vioResult{result}
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TimeTostring(times time.Time)string{
	//获取时间戳
	timestamp := times.Unix()
	//格式化为字符串,tm为Time类型
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 03:04:05")
}

func BooltoInt32(data bool) int32{
	if data {
		return 1
	}
	return 0
}

func Int32toBool(data int32) bool{
	if data>0 {
		return true
	}
	return false
}

func BooltoInt(data bool) int32{
	if data {
		return 1
	}
	return 0
}