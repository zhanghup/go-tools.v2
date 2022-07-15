package dm

import (
	"context"
	"github.com/zhanghup/go-tools.v2"
	"xorm.io/xorm"
)

type ISession[T any] interface {
	Id() string
	Context() context.Context
	Close() error

	Table(obj interface{}) ISession[T]
	SF(sqlstr string, querys ...any) ISession[T]
	Order(order ...string) ISession[T]

	Find() ([]T, error)
	FindBean(rowsSlicePtr interface{}) error
	Get() (T, bool, error)
	Exists() (v bool, err error)
	Count() (total int64, err error)
	Page(index, size int, count bool) (vs []T, v int, err error)

	Insert(bean ...T) error
	Update(bean any) error
	Delete() error
	Exec() error
}

type sessionEngine struct {
	id        string
	autoClose bool
	db        *xorm.Engine
	sess      *xorm.Session
	context   context.Context
}

type session[T any] struct {
	engine sessionEngine

	sfs *sessionSF[T]
}

func (s *session[T]) Table(obj interface{}) ISession[T] {
	tab := tools.RftTypeInfo(obj)
	s.sfs.tableName = s.engine.db.GetTableMapper().Obj2Table(tab.Name)
	return s
}

func (s *session[T]) Context() context.Context {
	if s.engine.context == nil {
		s.engine.context = context.Background()
	}
	return context.WithValue(s.engine.context, CONTEXT_SESSION, s.engine)
}

func (s *session[T]) Id() string {
	return s.engine.id
}

// Close 自动关闭session
func (s *session[T]) Close() error {
	if s.engine.sess.IsClosed() {
		return nil
	}
	return s.engine.sess.Close()
}

func (s *session[T]) _autoClose(fn func() error) error {
	err := fn()
	if err != nil {
		return err
	}

	if s.engine.autoClose {
		return s.Close()
	}
	return nil
}
