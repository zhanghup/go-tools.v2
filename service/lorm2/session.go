package lorm2

import (
	"context"
)

type ISession interface {
	Id() string
	SetId(id string)
	IsNew() bool
	Ctx() context.Context

	Table(bean any) ISession
	Find(bean any) error
	Get(bean any) (bool, error)

	Insert(bean ...any) error
	Update(bean any, condiBean ...any) error
	Delete(bean ...any) error
	Exec() error
	/*
		示例1：
			sql = "select * from user where a = ? and b = ?"
			querys = []interface{}{"a","b"}
		示例2：
			sql = "select * from user where a = :a and b = ?"
			querys = []interface{}{"b",map[string]interface{}{"a":"a"}}
		示例3：
			sql = "where a = ?"
			querys = []interface{}{"b"}
			bean = models.User
			>>> select user.* from user where a = ?
	*/
	SF(sql string, querys ...any) ISession
	Order(order ...string) ISession

	Page(index, size int, count bool, bean any) (int, error)
	Page2(index, size *int, count *bool, bean any) (int, error)
	Count() (int64, error)
	Int() (int, error)
	Int64() (int64, error)
	Float64() (float64, error)
	String() (string, error)
	Strings() ([]string, error)
	Exists() (bool, error)

	Map() ([]map[string]any, error)
	MapString() (v []map[string]string, err error)
}
