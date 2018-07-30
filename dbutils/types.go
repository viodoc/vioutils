package dbutil

type DB interface {
	preparer
	accesser
	Begin() Tx
	Close()
}

type Tx interface {
	preparer
	accesser
	Commit()
	Rollback()
}

type Stmt interface {
	Get(dest interface{}, args ...interface{}) error

	Select(dest interface{}, args ...interface{})

	Exec(args ...interface{}) Result
}

type NamedStmt interface {
	NamedGet(dest interface{}, arg interface{}) error

	NamedSelect(dest interface{}, arg interface{}) error

	NamedExec(arg interface{}) Result
}

type Result interface {
	LastInsertId() int64
	RowsAffected() int64
}

type accesser interface {
	NamedGet(dest interface{}, query string, arg interface{}) error

	Get(dest interface{}, query string, args ...interface{}) error

	NamedSelect(dest interface{}, query string, arg interface{}) error

	Select(dest interface{}, query string, args ...interface{}) error

	NamedExec(query string, arg interface{}) (Result,error)

	Exec(query string, args ...interface{}) Result
}

type preparer interface {
	Prepare(query string) Stmt

	PrepareNamed(query string) NamedStmt
}

