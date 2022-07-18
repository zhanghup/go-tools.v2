package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

// Slice 查找数据库对象,ctx可以为nil
func Slice[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) ([]Result, error) {
	return dm.Slice[Result](engine, ctx, beanKey, beanNameOrSql, field, param...)
}

func Slice2[Result any](ctx context.Context, beanKey, field string, param ...any) ([]Result, error) {
	tab := tools.RftTypeInfo(new(Result))
	tableName := Default().GetTableMapper().Obj2Table(tab.Name)
	return dm.Slice[Result](engine, ctx, beanKey, tableName, field, param...)
}

// Info 根据id查找数据库对象,ctx可以为nil
func Info[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) (*Result, error) {
	return dm.Info[Result](engine, ctx, beanKey, beanNameOrSql, field, param...)
}

func Info2[Result any](ctx context.Context, beanKey, field string, param ...any) (*Result, error) {
	tab := tools.RftTypeInfo(new(Result))
	tableName := Default().GetTableMapper().Obj2Table(tab.Name)

	return dm.Info[Result](engine, ctx, beanKey, tableName, field, param...)
}
