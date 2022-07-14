package db

import "github.com/zhanghup/go-tools.v2/service/dm"

func SF(sqlstr string, querys ...any) dm.ISession[any] {
	return dm.SF[any](engine, sqlstr, querys...)
}
