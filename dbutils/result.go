package dbutils

import "database/sql"

type vioResult struct {
	sql.Result
}

func (r *vioResult) LastInsertId() int64 {
	id, err := r.Result.LastInsertId()
	panicErr(err)
	return id
}

func (r *vioResult) RowsAffected() int64 {
	rows, err := r.Result.RowsAffected()
	panicErr(err)
	return rows
}

