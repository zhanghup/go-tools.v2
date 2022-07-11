package dm

import (
	_ "embed"
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/tog"
	"xorm.io/xorm"
	"xorm.io/xorm/log"
)

//go:embed config-default.yml
var defaultConfig []byte

type Config struct {
	Driver string `yaml:"driver"`
	Uri    string `yaml:"uri"`
	Debug  bool   `yaml:"debug"`
}

var engineMap = tools.NewCache[*xorm.Engine]()

const CONTEXT_SESSION = "context-xorm-session"

func InitXorm(ymlData ...[]byte) (*xorm.Engine, error) {
	cfg := struct {
		Db Config `json:"db" yaml:"db"`
	}{}

	err := tools.ConfOfByte(defaultConfig, &cfg)
	if err != nil {
		return nil, err
	}

	for _, data := range ymlData {
		if data == nil {
			continue
		}
		err = tools.ConfOfByte(data, &cfg)
		if err != nil {
			return nil, err
		}
	}

	return NewXorm(cfg.Db)
}

func NewXorm(cfg Config) (*xorm.Engine, error) {
	engineKey := fmt.Sprintf("%s___%s", cfg.Driver, cfg.Uri)

	v, ok := engineMap.Get(engineKey)
	if ok {
		return v, nil
	}

	engine, err := xorm.NewEngine(cfg.Driver, cfg.Uri)
	if err != nil {
		return nil, err
	}
	if cfg.Debug {
		engine.Logger().SetLevel(log.LOG_INFO)
		engine.SetLogger(log.NewSimpleLogger(tog.Writer))
		engine.ShowSQL(true)
	}
	engineMap.Set(engineKey, engine)
	return engine, err
}
