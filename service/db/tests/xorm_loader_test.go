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
		_, err := db.Info[User](context.Background(), fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestLoaderSlice(t *testing.T) {
	tools.Wait(10, func(nn int) {
		_, err := db.Slice[User](context.Background(), fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestLoaderInfoNil(t *testing.T) {
	tools.Wait(10, func(nn int) {
		_, err := db.Info[User](nil, fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestLoaderSliceNil(t *testing.T) {
	tools.Wait(10, func(nn int) {
		_, err := db.Slice[User](nil, fmt.Sprintf("%d", nn), "user", "username")
		if err != nil {
			t.Fatal(err)
		}
	})
}
