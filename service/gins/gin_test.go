package gins_test

import (
	"github.com/gin-gonic/gin"
	"github.com/zhanghup/go-tools.v2/service/gins"
	"testing"
)

func TestGin(t *testing.T) {
	gins.InitGin(func(g *gin.Engine) error {
		g.GET("a", func(c *gin.Context) {
			c.String(200, "a")
		})
		g.GET("b", func(c *gin.Context) {
			c.String(200, "b")
		})
		return nil
	})
}
