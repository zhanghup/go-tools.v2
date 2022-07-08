package dm

import (
	"xorm.io/xorm"
)

func SF[T any](db *xorm.Engine, sqlstr string, querys ...any) ISession[T] {
	return Session[T](db).SF(sqlstr, querys...)
}
