package tools_test

import (
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"testing"
)

func TestReverse(t *testing.T) {
	fmt.Println(tools.Reverse([]string{"1", "2", "3", "4"}))
}

func TestFmt(t *testing.T) {
	fmt.Println(tools.Fmt("<<n123>>", map[string]any{"n123": "你好"}))
	fmt.Println(tools.Fmt("<<   n123   >>", map[string]any{"n123": "你好"}))
	fmt.Println(tools.Fmt("<<n123abc>>", map[string]any{"n123abc": "你好"}))
	fmt.Println(tools.Fmt("<<n123abcABC>>", map[string]any{"n123abcABC": "你好"}))
	fmt.Println(tools.Fmt("<<n123abcABC你好你好>>", map[string]any{"n123abcABC你好你好": "你好"}))
	fmt.Println(tools.Fmt("<<n123abcABC你好你好>> vvvvvv <<a>>", map[string]any{"n123abcABC你好你好": "你好", "a": "哈哈哈哈哈"}))
	fmt.Println(tools.Fmt("%s 的哈哈的发挥 <<n123>>", "咳咳咳咳咳", map[string]any{"n123": "你好"}))
	fmt.Println(tools.Fmt("的哈哈的发挥 <<n123>>  %s %d", map[string]any{"n123": "你好"}, "咳咳咳咳咳", 2222))
	fmt.Println(tools.Fmt("的哈哈的发挥 <<n123>>  %s %d {{.kkk}}", map[string]any{"n123": "你好", "kkk": "觉得咖啡机"}, "咳咳咳咳咳", 2222))
}
