package lorm

import (
	"context"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

func Session[T any](db *xorm.Engine) *session[T] {
	newSession := &session[T]{
		engine: sessionEngine{
			id:        tools.UUID(),
			db:        db,
			sess:      db.NewSession(),
			autoClose: true,
		},
	}

	newSession.context = context.WithValue(context.Background(), CONTEXT_SESSION, newSession.engine)
	return newSession
}

func Context[T any](db *xorm.Engine, ctx context.Context) *session[T] {
	v := ctx.Value(CONTEXT_SESSION)
	if v != nil {
		sessOld, ok := v.(sessionEngine)
		if ok && !sessOld.sess.IsClosed() {
			return &session[T]{context: ctx, engine: sessOld}
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

	newSession.context = context.WithValue(ctx, CONTEXT_SESSION, newSession.engine)
	return newSession
}

func TS(ctx context.Context, db *xorm.Engine, fn func(ctx context.Context) error) error {
	commit := ctx.Value(CONTEXT_SESSION)
	sess := Context[int](db, ctx)
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
