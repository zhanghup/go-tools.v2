package tests

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestFind(t *testing.T) {
	users := []User{}
	err := db.SF("age = 0").FindBean(&users)
	if err != nil {
		t.Fatal(err)
	}
}
