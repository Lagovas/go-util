package udb

import (
	"database/sql"
)

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type SqlCursor struct {
	cols       []string
	vals, ptrs []interface{}
}

func Exec(execer Execer, isInsert bool, query string, args ...interface{}) (result int64, err error) {
	var res sql.Result
	if res, err = execer.Exec(query, args...); err == nil {
		if isInsert {
			result, err = res.LastInsertId()
		} else {
			result, err = res.RowsAffected()
		}
	}
	return
}

func (me *SqlCursor) PrepareColumns(cursor *sql.Rows) (err error) {
	if me.cols, err = cursor.Columns(); err == nil {
		me.vals, me.ptrs = make([]interface{}, len(me.cols), len(me.cols)), make([]interface{}, len(me.cols), len(me.cols))
		for i := 0; i < len(me.ptrs); i++ {
			me.ptrs[i] = &me.vals[i]
		}
	}
	return
}

func (me *SqlCursor) Scan(cursor *sql.Rows) (rec map[string]interface{}, err error) {
	if err = cursor.Scan(me.ptrs...); err == nil {
		rec = make(map[string]interface{}, len(me.vals))
		var bytes []byte
		var ok bool
		for i := 0; i < len(me.vals); i++ {
			if bytes, ok = me.vals[i].([]byte); ok {
				me.vals[i] = string(bytes)
			}
			rec[me.cols[i]] = me.vals[i]
		}
	}
	return
}
