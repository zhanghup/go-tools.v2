package dm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

func Session[T any](db *xorm.Engine) ISession[T] {

	newSession := &session[T]{
		engine: sessionEngine{
			id:        tools.UUID(),
			db:        db,
			sess:      db.NewSession(),
			autoClose: true,
		},
	}

	newSession.engine.context = context.WithValue(context.Background(), CONTEXT_SESSION, newSession.engine)
	newSession.sfs = newSessionSf[T](db, newSession.engine.context)
	return newSession
}

func Context[T any](db *xorm.Engine, ctx context.Context) ISession[T] {
	if ctx == nil {
		ctx = context.Background()
	}

	v := ctx.Value(CONTEXT_SESSION)
	if v != nil {
		sessOld, ok := v.(sessionEngine)
		if ok && !sessOld.sess.IsClosed() {
			sessOld.context = ctx
			return &session[T]{engine: sessOld, sfs: newSessionSf[T](db, ctx)}
		}
	}

	newSession := &session[T]{
		engine: sessionEngine{
			id:        tools.UUID(),
			db:        db,
			sess:      db.NewSession(),
			autoClose: false,
		},
	}

	newSession.engine.context = context.WithValue(ctx, CONTEXT_SESSION, newSession.engine)
	newSession.sfs = newSessionSf[T](db, newSession.engine.context)
	return newSession
}

func TS(ctx context.Context, db *xorm.Engine, fn func(ctx context.Context) error) error {
	if ctx == nil {
		ctx = context.Background()
	}

	commit := ctx.Value(CONTEXT_SESSION)
	sess := Context[any](db, ctx).(*session[any])
	if commit == nil {
		_ = sess.engine.sess.Begin()
	}
	err := fn(sess.Context())
	if err != nil {
		_ = sess.engine.sess.Rollback()
		return err
	}
	if commit == nil {
		_ = sess.engine.sess.Commit()
		return sess.Close()
	}
	return nil

}
