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

	_, _, err := db.SF(`select * from {{ withs "age" }}`).Get()
	if err != nil {
		t.Fatal(err)
	}
}
