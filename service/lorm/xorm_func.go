package lorm

import (
	"xorm.io/xorm"
)

func Insert(db *xorm.Engine, bean ...any) error {
	return Session[any](db).Insert(bean...)
}

func Update(db *xorm.Engine, bean any, condiBean ...any) error {
	return Session[any](db).Update(bean, condiBean...)
}

func Delete(db *xorm.Engine, bean ...any) error {
	return Session[any](db).Delete(bean...)
}

func Exec(db *xorm.Engine) error {
	return Session[any](db).Exec()
}
