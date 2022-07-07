package lorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

type sessionEngine struct {
	id        string
	autoClose bool
	db        *xorm.Engine
	sess      *xorm.Session
}

type session[T any] struct {
	context context.Context
	engine  sessionEngine

	tableName string
	sql       string
	query     map[string]any
	args      []any

	withs   []string
	orderby []string
}

func (s *session[T]) Context() context.Context {
	if s.context == nil {
		s.context = context.Background()
	}
	return context.WithValue(s.context, CONTEXT_SESSION, s.engine)
}

func (s *session[T]) Id() string {
	return s.engine.id
}

func (s *session[T]) _autoClose(fn func() error) error {
	err := fn()
	if err != nil {
		return err
	}
	s.tableName = ""

	if s.engine.autoClose {
		return s.Close()
	}
	return nil
}

// Close 自动关闭session
func (s *session[T]) Close() error {
	if s.engine.sess.IsClosed() {
		return nil
	}
	return s.engine.sess.Close()
}

func (s *session[T]) Table(bean any) *session[T] {
	if s.tableName != "" {
		return s
	}

	switch bean.(type) {
	case string:
		s.tableName = bean.(string)
	case *string:
		s.tableName = *(bean.(*string))
	default:
		tab := tools.RftTypeInfo(bean)
		s.tableName = s.engine.db.GetTableMapper().Obj2Table(tab.Name)
	}

	return s
}
