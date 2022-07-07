package lorm

import (
	"strings"
)

func (s *session[T]) Insert(bean ...any) error {
	return s._autoClose(func() error {
		if s.tableName == "" {
			s.Table(bean)
		}
		_, err := s.engine.sess.Table(s.tableName).Insert(bean...)
		return err
	})
}

func (s *session[T]) Update(bean any, condiBean ...any) error {
	return s._autoClose(func() error {
		if s.tableName == "" {
			s.Table(bean)
		}

		sqlstr := strings.TrimSpace(s._sql(false))
		_, err := s.engine.sess.Table(s.tableName).Where(sqlstr, s.args...).Update(bean, condiBean...)
		return err
	})
}

func (s *session[T]) Delete(bean ...any) error {
	return s._autoClose(func() error {
		if s.tableName == "" {
			s.Table(bean)
		}

		sqlstr := strings.TrimSpace(s._sql(false))
		_, err := s.engine.sess.Table(s.tableName).Where(sqlstr, s.args...).Delete(bean...)
		return err
	})
}

func (s *session[T]) Exec() error {
	return s._autoClose(func() error {
		sqls := []any{s._sql_with() + " " + s._sql(true)}
		_, err := s.engine.sess.Exec(append(sqls, s.args...)...)
		return err
	})
}
