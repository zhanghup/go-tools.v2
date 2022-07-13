package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

// Slice 查找数据库对象,ctx可以为nil
func Slice[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) ([]Result, error) {
	return dm.Slice[Result](engine, ctx, beanKey, beanNameOrSql, field, param...)
}

func SliceId[Result any](ctx context.Context, beanKey, beanNameOrSql string, param ...any) ([]Result, error) {
	return dm.SliceId[Result](engine, ctx, beanKey, beanNameOrSql, param...)
}

// Info 根据id查找数据库对象,ctx可以为nil
func Info[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) (*Result, error) {
	return dm.Info[Result](engine, ctx, beanKey, beanNameOrSql, field, param...)
}

func InfoId[Result any](ctx context.Context, beanKey, beanNameOrSql string, param ...any) (*Result, error) {
	return dm.InfoId[Result](engine, ctx, beanKey, beanNameOrSql, param...)
}
