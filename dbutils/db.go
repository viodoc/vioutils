package dbutils

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

type vioDB struct {
	*sqlx.DB
}
var (
	viodb *vioDB
	once sync.Once
)

func NewDB(driverName, dsn string) DB {
	once.Do(func() {
		viodb = &vioDB{}
		viodb.DB = sqlx.MustOpen(driverName, dsn)
		viodb.DB.SetMaxIdleConns(10)
		viodb.DB.SetMaxOpenConns(500)
	})
	return viodb
}

func (db *vioDB) Get(dest interface{}, query string, args ...interface{}) error {
	return get(db.DB, dest, query, args...)
}

func (db *vioDB) NamedGet(dest interface{}, query string, arg interface{}) error {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	return namedGet(stmt, dest, arg)
}

func (db *vioDB) NamedSelect(dest interface{}, query string, arg interface{}) {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	namedSelect(stmt, dest, arg)
}

func (db *vioDB) Select(dest interface{}, query string, args ...interface{}) {
	selectx(db.DB, dest, query, args...)
}

func (db *vioDB) NamedExec(query string, arg interface{}) Result {
	return namedExec(db.DB, query, arg)
}

func (db *vioDB) Exec(query string, args ...interface{}) Result {
	return exec(db.DB, query, args...)
}

func (db *vioDB) Begin() Tx {
	tx := db.DB.MustBegin()
	return &vioTx{Tx: tx}
}

func (db *vioDB) Prepare(query string) Stmt {
	stmt, err := db.DB.Preparex(query)
	panicErr(err)
	return &vioStmt{Stmt: stmt}
}

func (db *vioDB) PrepareNamed(query string) NamedStmt {
	stmt, err := db.DB.PrepareNamed(query)
	panicErr(err)
	return &vioNamedStmt{NamedStmt: stmt}
}

func (db *vioDB) Close() {
	err:=db.DB.Close()
	panicErr(err)
}
