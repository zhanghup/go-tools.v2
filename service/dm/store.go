package dm

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"xorm.io/xorm"
)

var store = tools.NewCache[map[string]any]()

func dbKey(db *xorm.Engine) string {
	return fmt.Sprintf("%s___%s", db.DriverName(), db.DataSourceName())
}

func templateFunctions(db *xorm.Engine, startWith string) map[string]any {
	key := startWith + dbKey(db)

	v, ok := store.Get(key)
	if !ok {
		return map[string]any{}
	}
	return v
}

func SetWiths(db *xorm.Engine, name string, f func(ctx context.Context) string) {
	key := "___with_" + dbKey(db)
	kv := fmt.Sprintf("___with_%s", name)

	v, ok := store.Get(key)
	if !ok {
		store.Set(key, map[string]any{kv: f})
		return
	}
	v[kv] = f
	store.Set(key, v)
}

func SetTemplates(db *xorm.Engine, name string, f any) {
	key := "___templates_" + dbKey(db)
	kv := fmt.Sprintf("___templates_%s", name)

	v, ok := store.Get(key)
	if !ok {
		store.Set(key, map[string]any{kv: f})
		return
	}
	v[kv] = f
	store.Set(key, v)
}

func SetContexts(db *xorm.Engine, name string, f func(ctx context.Context) string) {
	key := "___contexts_" + dbKey(db)
	kv := fmt.Sprintf("___contexts_%s", name)

	v, ok := store.Get(key)
	if !ok {
		store.Set(key, map[string]any{kv: f})
		return
	}
	v[kv] = f
	store.Set(key, v)
}
