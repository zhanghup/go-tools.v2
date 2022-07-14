package tests

import (
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestInsert(t *testing.T) {

	for i := 0; i < 10; i++ {
		err := db.Session[User]().Insert(User{
			Id:   tools.UUID(),
			Name: "zander",
			Age:  i,

			Username: fmt.Sprintf("%d", i),
			Password: "Aa123456.",
		})

		if err != nil {
			t.Fatal(err)
		}
	}

}

func TestUpdate(t *testing.T) {
	err := db.SF("username = ?", "zander").Table(User{}).Update(map[string]any{
		"age": 1,
	})
	if err != nil {
		t.Fatal(err)
	}
}
