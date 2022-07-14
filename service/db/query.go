package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

func SF(ctx context.Context, sqlstr string, querys ...any) dm.ISession[any] {
	return dm.SF[any](ctx, engine, sqlstr, querys...)
}

func SFC[T any](ctx context.Context, sqlstr string, querys ...any) dm.ISession[T] {
	return dm.SF[T](ctx, engine, sqlstr, querys...)
}
