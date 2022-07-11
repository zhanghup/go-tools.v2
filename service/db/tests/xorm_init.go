package tests

import (
	"github.com/zhanghup/go-tools.v2/service/db"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`

	Username string `json:"username"`
	Password string `json:"password"`
}

func init() {
	err := db.Init(db.Config{
		Uri:    "./data.db",
		Driver: "sqlite3",
		Debug:  true,
	})
	if err != nil {
		panic(err)
	}
}
