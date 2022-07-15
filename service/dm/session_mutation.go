package dm

import "github.com/zhanghup/go-tools.v2"

func (s *session[T]) Insert(bean ...T) error {
	return s._autoClose(func() error {
		_, err := s.engine.sess.Table(s.sfs.tableName).Insert(tools.AnyToAny(bean...)...)
		return err
	})
}

func (s *session[T]) Update(bean any) error {
	return s._autoClose(func() error {

		_, err := s.engine.sess.Table(s.sfs.tableName).Where(s.sfs.SQL(false, false), s.sfs.sqlArgs...).Update(bean)
		return err
	})
}

func (s *session[T]) Delete() error {
	return s._autoClose(func() error {
		_, err := s.engine.sess.Table(s.sfs.tableName).Where(s.sfs.SQL(false, false), s.sfs.sqlArgs...).Delete()
		return err
	})
}

func (s *session[T]) Exec() error {
	return s._autoClose(func() error {
		_, err := s.engine.sess.Exec(append([]any{s.sfs.SQL(false, false)}, s.sfs.sqlArgs...)...)
		return err
	})
}
