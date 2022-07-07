package lorm

import (
	_ "github.com/mattn/go-sqlite3"
	"testing"
	"xorm.io/xorm"
)

var engine *xorm.Engine

func TestStore(t *testing.T) {
	getWiths(engine)
}

func init() {
	e, err := NewXorm(Config{
		Uri:    "./data.db",
		Driver: "sqlite3",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
	engine = e
}
