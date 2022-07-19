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

func Find[T any](ctx context.Context, sqlOrArgs ...any) ([]T, error) {
	return dm.Find[T](ctx, engine, sqlOrArgs...)
}

func Get[T any](ctx context.Context, sqlOrArgs ...any) (*T, error) {
	return dm.Get[T](ctx, engine, sqlOrArgs...)
}

func Exists[T any](ctx context.Context, sqlOrArgs ...any) (bool, error) {
	return dm.Exists[T](ctx, engine, sqlOrArgs...)
}
