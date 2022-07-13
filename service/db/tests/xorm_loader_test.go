package tests

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/db"
	"testing"
)

func TestLoaderInfo(t *testing.T) {
	tools.Wait(10, func(nn int) {
		res, err := db.Info[User](context.Background(), fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	})
}

func TestLoaderSlice(t *testing.T) {
	tools.Wait(10, func(nn int) {
		res, err := db.Slice[User](context.Background(), fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			panic(err)
		}
		fmt.Println(res)
	})
}
