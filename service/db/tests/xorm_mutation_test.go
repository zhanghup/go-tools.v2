package tests

import (
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestInsert(t *testing.T) {

}

func TestUpdate(t *testing.T) {
	err := db.SF(nil, "username = ?", "zander").Table(User{}).Update(map[string]any{
		"age": 1,
	})
	if err != nil {
		t.Fatal(err)
	}
}
