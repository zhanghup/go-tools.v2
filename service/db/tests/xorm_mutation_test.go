package tests

import (
	"github.com/zhanghup/go-tools"
	"github.com/zhanghup/go-tools/service/db"
	"testing"
)

func TestInsert(t *testing.T) {
	err := db.Session[User]().Insert(User{
		Id:   tools.UUID(),
		Name: "zander",
		Age:  10,

		Username: "zander",
		Password: "Aa123456.",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdate(t *testing.T) {
	err := db.SF[User]("username = ?", "zander").Update(map[string]any{
		"age": 1,
	})
	if err != nil {
		t.Fatal(err)
	}
}
