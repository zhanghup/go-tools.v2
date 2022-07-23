package tests

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestWith(t *testing.T) {
	db.SetWiths("age", func(ctx context.Context) string {
		return "select 1"
	})

	_, err := db.SF(nil, `select * from {{ withs "age" }}`).Get()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSF(t *testing.T) {
	_, err := db.SF(nil, "select 1 from user limit 1").Get()
	if err != nil {
		t.Fatal(err)
	}
}

func TestSFC(t *testing.T) {
	v, err := db.SFC[int](nil, "select 1 from user limit 1").GetOne()
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

	v2, err := db.SFC[User](nil, "id = ?", "123").GetOne()
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
	user, err := db.Get[User](nil)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal("查询异常")
	}
	if user.Id != "111" {
		t.Fatal("查询异常[4]")
	}

	user, err = db.Get[User](nil, "age = ?", 4)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
		t.Fatal("查询异常[2]")
	}

	user, err = db.Get[User](nil, "age = ? limit 10", 4)
	if err != nil {
		t.Fatal(err)
	}
	if user == nil {
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

func TestLike(t *testing.T) {
	users, err := db.Find[User](nil, "name like:?", []int{3, 5}, "zander")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}

func TestBetween(t *testing.T) {
	users, err := db.Find[User](nil, "age between:? and name like:?", []int{3, 5}, "zander")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}

func TestIn(t *testing.T) {
	users, err := db.Find[User](nil, "age in:?", []int{3, 5})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}
