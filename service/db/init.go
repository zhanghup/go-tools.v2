package db

import (
	"github.com/zhanghup/go-tools.v2/service/dm"
	"xorm.io/xorm"
)

var engine *xorm.Engine

type Config dm.Config

func Default() *xorm.Engine {
	if engine == nil {
		panic("默认数据库未初始化")
	}
	return engine
}

func InitXorm(ymlData ...[]byte) error {
	db, err := dm.InitXorm(ymlData...)
	if err != nil {
		return err
	}
	engine = db
	return nil
}

func Init(cfg Config) error {
	db, err := dm.NewXorm(dm.Config(cfg))
	if err != nil {
		return err
	}
	engine = db
	return nil
}
