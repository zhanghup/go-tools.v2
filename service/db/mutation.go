package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

func Insert[T any](ctx context.Context, bean T) error {
	return dm.Insert[T](ctx, engine, bean)
}
