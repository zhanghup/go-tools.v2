package db

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zhanghup/go-tools.v2/service/dm"
)

func Session[T any](ctx context.Context) dm.ISession[T] {
	return dm.Session[T](engine, ctx)
}

func Context[T any](ctx context.Context) dm.ISession[T] {
	return dm.Context[T](engine, ctx)
}

/*
	TemplateFuncWith Sql With 模板，

	初始化模板：
	db.TemplateFunc("users",function(ctx context.Context) string{
		return "select id,name from user"
	})

	select * from {{ withs "users" }}
	=>
	{{ sql_with_users }} select * from __sql_with_users
	=>
	with recursive _ as (select 1),__sql_with_users as (select id,name from user) select * from __sql_with_users
*/
func SetWiths(name string, f func(ctx context.Context) string) {
	dm.SetWiths(engine, name, f)
}

func SetTemplates(name string, f any) {
	dm.SetTemplates(engine, name, f)
}

func SetContexts(name string, f func(ctx context.Context) string) {
	dm.SetContexts(engine, name, f)
}

func TS(ctx context.Context, fn func(ctx context.Context) error) error {
	return dm.TS(ctx, engine, fn)
}
