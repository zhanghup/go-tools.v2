package tools

import "github.com/zhanghup/go-tools/lib/pinyin"

// Py 中文转拼音，例如：“你好” => “nh”
func Py(str string) string {
	return pinyin.ExportPy(str)
}

// PY 中文转拼音，例如：“你好” => “NH”
func PY(str string) string {
	return pinyin.ExportPY(str)
}

// Pinyin 中文转拼音，例如：“你好” => “nihao”
func Pinyin(str string) string {
	return pinyin.ExportPinyin(str)
}

// PINYIN 中文转拼音，例如：“你好” => “NIHAO”
func PINYIN(str string) string {
	return pinyin.ExportPINYIN(str)
}
