package db

import "github.com/zhanghup/go-tools.v2/service/dm"

func SF[T any](sqlstr string, querys ...any) dm.ISession[T] {
	return dm.SF[T](engine, sqlstr, querys...)
}
