package dbutils

import (
	"github.com/jmoiron/sqlx"
)

type vioStmt struct {
	*sqlx.Stmt
}

func (vs *vioStmt) Get(dest interface{}, args ...interface{}) error {
	return vs.Stmt.Get(dest, args...)
}

func (vs *vioStmt) Select(dest interface{}, args ...interface{}) {
	err := vs.Stmt.Select(dest, args...)
	panicErr(err)
}

func (vs *vioStmt) Exec(args ...interface{}) Result {
	result := vs.Stmt.MustExec(args...)
	return &vioResult{result}
}

type vioNamedStmt struct {
	*sqlx.NamedStmt
}

func (vns *vioNamedStmt) NamedGet(dest interface{}, arg interface{}) error {
	return namedGet(vns.NamedStmt, dest, arg)
}

func (vns *vioNamedStmt) NamedSelect(dest interface{}, arg interface{}) {
	namedSelect(vns.NamedStmt, dest, arg)
}
func (vns *vioNamedStmt) NamedExec(arg interface{}) Result {
	result := vns.NamedStmt.MustExec(arg)
	return &vioResult{result}
}