package tests

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools/service/db"
	"testing"
)

func TestWith(t *testing.T) {
	db.SetWiths("age", func(ctx context.Context) string {
		return "select 1"
	})

	users, _, err := db.SF[int](`select * from {{ withs "age" }}`).Get()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(users)
}
