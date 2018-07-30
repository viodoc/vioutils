package dbutil

import "github.com/jmoiron/sqlx"

type vioTx struct {
	*sqlx.Tx
}

func (tx *vioTx) NamedGet(dest interface{}, query string, arg interface{}) error {
	stmt, err := tx.Tx.PrepareNamed(query)
	panicErr(err)
	return namedGet(stmt, dest, arg)
}

func (tx *vioTx) Get(dest interface{}, query string, args ...interface{}) error {
	return get(tx.Tx, dest, query, args...)
}

func (tx *vioTx) NamedSelect(dest interface{}, query string, arg interface{}) error {
	stmt, err := tx.Tx.PrepareNamed(query)
	if err != nil{
		return err
	}
	return namedSelect(stmt, dest, arg)
}

func (tx *vioTx) Select(dest interface{}, query string, args ...interface{}) error {
	return selectx(tx.Tx, dest, query, args...)
}

func (tx *vioTx) NamedExec(query string, arg interface{}) (Result,error) {
	return namedExec(tx.Tx, query, arg)
}

func (tx *vioTx) Exec(query string, args ...interface{}) Result {
	return exec(tx.Tx, query, args...)
}

func (tx *vioTx) Commit() {
	err := tx.Tx.Commit()
	panicErr(err)
}

func (tx *vioTx) Rollback() {
	err := tx.Tx.Rollback()
	panicErr(err)
}

func (tx *vioTx) Prepare(query string) Stmt {
	stmt, err := tx.Tx.Preparex(query)
	panicErr(err)
	return &vioStmt{Stmt: stmt}
}

func (db *vioTx) PrepareNamed(query string) NamedStmt {
	stmt, err := db.Tx.PrepareNamed(query)
	panicErr(err)
	return &vioNamedStmt{NamedStmt: stmt}
}

