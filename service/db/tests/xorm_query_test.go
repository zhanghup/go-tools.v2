package tests

import (
	"context"
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestWith(t *testing.T) {
	db.SetWiths("age", func(ctx context.Context) string {
		return "select 1"
	})

	_, _, err := db.SF(nil, `select * from {{ withs "age" }}`).Get()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSF(t *testing.T) {
	_, _, err := db.SF(nil, "select 1 from user limit 1").Get()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSFC(t *testing.T) {
	v, _, err := db.SFC[int](nil, "select 1 from user limit 1").Get()
	if err != nil {
		t.Fatal(err)
	}
	if v != 1 {
		t.Fatal("查询失败")
	}

	err = db.Insert(nil, User{Id: "123"})
	if err != nil {
		t.Fatal()
	}

	v2, _, err := db.SFC[User](nil, "id = ?", "123").Get()
	if err != nil {
		t.Fatal(err)
	}
	if v2.Id != "123" {
		t.Fatal("查询失败2")
	}
}

func TestFind(t *testing.T) {
	users, err := db.Find[User](nil)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 10 {
		t.Fatal("查询异常")
	}

	users, err = db.Find[User](nil, "age = ?", 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Fatal("查询异常[2]")
	}

	users, err = db.Find[User](nil, "age = ? limit 10", 4)
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 1 {
		t.Fatal("查询异常[3]")
	}
}

func TestGet(t *testing.T) {
	user, ok, err := db.Get[User](nil)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常")
	}
	if user.Id != "111" {
		t.Fatal("查询异常[4]")
	}

	user, ok, err = db.Get[User](nil, "age = ?", 4)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常[2]")
	}

	user, ok, err = db.Get[User](nil, "age = ? limit 10", 4)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常[3]")
	}
}

func TestExists(t *testing.T) {
	ok, err := db.Exists[User](nil)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常")
	}

	ok, err = db.Exists[User](nil, "age = ?", 4)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常[2]")
	}

	ok, err = db.Exists[User](nil, "age = ? limit 10", 4)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("查询异常[3]")
	}
}
