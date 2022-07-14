package tests

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools.v2"
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
	db.Default().DropTables(User{})
	db.Default().Sync2(User{})

	err = db.Insert(nil, User{
		Id:   "111",
		Name: "zander",
		Age:  999,

		Username: fmt.Sprintf("%d", 888),
		Password: "Aa123456.",
	})

	if err != nil {
		panic(err)
	}

	for i := 0; i < 9; i++ {
		err := db.Session[User]().Insert(User{
			Id:   tools.UUID(),
			Name: "zander",
			Age:  i,

			Username: fmt.Sprintf("%d", i),
			Password: "Aa123456.",
		})

		if err != nil {
			panic(err)
		}
	}
}
