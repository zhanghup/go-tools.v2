package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

func Insert[T any](ctx context.Context, bean ...T) error {
	return dm.Insert[T](ctx, engine, bean...)
}

func Update[T any](ctx context.Context, sqlOrArgsOrBean ...any) error {
	return dm.Update[T](ctx, engine, sqlOrArgsOrBean...)
}

func Delete[T any](ctx context.Context, sqlOrArgsOrBean ...any) error {
	return dm.Delete[T](ctx, engine, sqlOrArgsOrBean...)
}

func Exec[T any](ctx context.Context, sqlOrArgs ...any) error {
	return dm.Exec[T](ctx, engine, sqlOrArgs...)
}
