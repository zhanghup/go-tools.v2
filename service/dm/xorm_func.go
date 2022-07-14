package dm

import (
	"context"
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

func Find[T any](ctx context.Context, db *xorm.Engine) ([]T, error) {
	return inSession[T](ctx, db).Find()
}
