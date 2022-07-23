package dm

import (
	"context"
	"github.com/zhanghup/go-tools.v2"
	"xorm.io/xorm"
)

func Session[T any](db *xorm.Engine, ctx context.Context) ISession[T] {

	c := context.Background()
	if ctx != nil {
		c = ctx
	}

	newSession := &session[T]{
		engine: sessionEngine{
			id:        tools.UUID(),
			db:        db,
			sess:      db.NewSession(),
			autoClose: true,
		},
	}

	newSession.engine.context = context.WithValue(c, CONTEXT_SESSION, newSession.engine)
	newSession.sfs = newSessionSf[T](db, newSession.engine.context)
	return newSession
}

func Context[T any](db *xorm.Engine, ctx context.Context) ISession[T] {
	if ctx == nil {
		ctx = context.Background()
	}

	if v, ok := hasSession(ctx); ok {
		return &session[T]{engine: v, sfs: newSessionSf[T](db, ctx)}
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
