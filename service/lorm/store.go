package lorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"xorm.io/xorm"
)

var store = tools.NewCache[map[string]any]()

func dbKey(db *xorm.Engine) string {
	return fmt.Sprintf("%s___%s", db.DriverName(), db.DataSourceName())
}

func getWiths(db *xorm.Engine) map[string]any {
	key := dbKey(db) + "___with"

	v, ok := store.Get(key)
	if !ok {
		return map[string]any{}
	}
	return v
}

func SetWiths(db *xorm.Engine, name string, f any) {
	key := dbKey(db) + "___with"

	v, ok := store.Get(key)
	if !ok {
		store.Set(key, map[string]any{name: f})
		return
	}
	v[name] = f
	store.Set(key, v)
}

func getTemplates(db *xorm.Engine) map[string]any {
	key := dbKey(db) + "___templates"

	v, ok := store.Get(key)
	if !ok {
		return map[string]any{}
	}
	return v
}

func SetTemplates(db *xorm.Engine, name string, f any) {
	key := dbKey(db) + "___templates"

	v, ok := store.Get(key)
	if !ok {
		store.Set(key, map[string]any{name: f})
		return
	}
	v[name] = f
	store.Set(key, v)
}
