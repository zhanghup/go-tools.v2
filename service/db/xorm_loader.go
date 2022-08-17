package db

import (
	"context"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

// Slice2 查找数据库对象,ctx可以为nil
func Slice2[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) ([]Result, error) {
	return dm.Slice[Result](engine, ctx, beanKey, beanNameOrSql, field, []string{}, param...)
}

/*
Slice

db.Slice[beans.User](ctx,"62ecc37312dc3d00012bff2a","id",[]string{})
db.Slice[beans.User](ctx,"62ecc37312dc3d00012bff2a",`id|corp = {{ ctx "corp" }}`,[]string{})  >> "|"后面代表在where后面添加的查询条件
db.Slice[beans.User](ctx,"62ecc37312dc3d00012bff2a",`id|corp = ?`,[]string{},"0000")  >> "|"后面代表在where后面添加的查询条件
*/
func Slice[Result any](ctx context.Context, beanKey, field string, order []string, param ...any) ([]Result, error) {
	tab := tools.RftTypeInfo(new(Result))
	tableName := Default().GetTableMapper().Obj2Table(tab.Name)
	return dm.Slice[Result](engine, ctx, beanKey, tableName, field, order, param...)
}

// Info2 根据id查找数据库对象,ctx可以为nil
func Info2[Result any](ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) (*Result, error) {
	return dm.Info[Result](engine, ctx, beanKey, beanNameOrSql, field, param...)
}

/*
Info

db.Info[beans.User](ctx,"62ecc37312dc3d00012bff2a","id")
db.Info[beans.User](ctx,"62ecc37312dc3d00012bff2a",`id|corp = {{ ctx "corp" }}`)  >> "|"后面代表在where后面添加的查询条件
db.Info[beans.User](ctx,"62ecc37312dc3d00012bff2a",`id|corp = ?`,"0000")  >> "|"后面代表在where后面添加的查询条件
*/
func Info[Result any](ctx context.Context, beanKey, field string, param ...any) (*Result, error) {
	tab := tools.RftTypeInfo(new(Result))
	tableName := Default().GetTableMapper().Obj2Table(tab.Name)

	return dm.Info[Result](engine, ctx, beanKey, tableName, field, param...)
}
