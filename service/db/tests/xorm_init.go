package tests

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools.v2/service/db"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`

	Username string `json:"username"`
	Password string `json:"password"`
}
type UserDep struct {
	Id  string `json:"id"`
	Uid string `json:"uid"`
	Dep string `json:"dep"`
}

type Dep struct {
	Id   string `json:"id"`
	Name string `json:"name"`
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
	db.Default().DropTables(User{}, UserDep{}, Dep{})
	db.Default().Sync2(User{}, UserDep{}, Dep{})

	err = db.Insert(nil, User{
		Id:   "1",
		Name: "张三",
		Age:  18,

		Username: fmt.Sprintf("%d", 888),
		Password: "Aa123456.",
	}, User{
		Id:   "2",
		Name: "李四",
		Age:  19,

		Username: fmt.Sprintf("%d", 888),
		Password: "Aa123456.",
	}, User{
		Id:       "3",
		Name:     "王五",
		Age:      20,
		Username: fmt.Sprintf("%d", 888),
		Password: "Aa123456.",
	})
	if err != nil {
		panic(err)
	}

	err = db.Insert(nil, Dep{
		Id:   "1",
		Name: "A部门",
	}, Dep{
		Id:   "2",
		Name: "B部门",
	}, Dep{
		Id:   "3",
		Name: "C部门",
	})
	if err != nil {
		panic(err)
	}

	err = db.Insert(nil, UserDep{
		Id:  "1",
		Uid: "1",
		Dep: "1",
	}, UserDep{
		Id:  "2",
		Uid: "1",
		Dep: "2",
	}, UserDep{
		Id:  "3",
		Uid: "2",
		Dep: "1",
	}, UserDep{
		Id:  "4",
		Uid: "3",
		Dep: "3",
	})
	if err != nil {
		panic(err)
	}
}
