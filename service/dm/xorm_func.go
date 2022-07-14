package dm

import (
	"context"
	"errors"
	"xorm.io/xorm"
)

func hasSession(ctx context.Context) (sessionEngine, bool) {
	v := ctx.Value(CONTEXT_SESSION)
	if v != nil {
		sessOld, ok := v.(sessionEngine)
		if ok && !sessOld.sess.IsClosed() {
			sessOld.context = ctx
			return sessOld, true
		}
	}

	return sessionEngine{}, false
}

func inSession[T any](ctx context.Context, db *xorm.Engine) ISession[T] {
	var ss ISession[T]

	if _, ok := hasSession(ctx); ok {
		ss = Context[T](db, ctx)
	} else {
		ss = Session[T](db)
	}

	return ss
}

func SF[T any](ctx context.Context, db *xorm.Engine, sqlstr string, querys ...any) ISession[T] {
	return inSession[T](ctx, db).SF(sqlstr, querys...)
}

func Insert[T any](ctx context.Context, db *xorm.Engine, bean T) error {
	return inSession[T](ctx, db).Insert(bean)
}

func Find[T any](ctx context.Context, db *xorm.Engine, sqlOrArgs ...any) ([]T, error) {
	s := inSession[T](ctx, db)
	if len(sqlOrArgs) > 0 {
		switch sqlOrArgs[0].(type) {
		case string:
			return s.SF(sqlOrArgs[0].(string), sqlOrArgs[1:]...).Find()
		default:
			return nil, errors.New("sqlOrArgs异常")
		}
	}

	return s.Find()
}

func Get[T any](ctx context.Context, db *xorm.Engine, sqlOrArgs ...any) (T, bool, error) {
	s := inSession[T](ctx, db)
	if len(sqlOrArgs) > 0 {
		switch sqlOrArgs[0].(type) {
		case string:
			return s.SF(sqlOrArgs[0].(string), sqlOrArgs[1:]...).Get()
		default:
			return nil, false, errors.New("sqlOrArgs异常")
		}
	}

	return s.Get()
}

func Exists[T any](ctx context.Context, db *xorm.Engine, sqlOrArgs ...any) (bool, error) {
	s := inSession[T](ctx, db)
	if len(sqlOrArgs) > 0 {
		switch sqlOrArgs[0].(type) {
		case string:
			return s.SF(sqlOrArgs[0].(string), sqlOrArgs[1:]...).Exists()
		default:
			return false, errors.New("sqlOrArgs异常")
		}
	}

	return s.Exists()
}
