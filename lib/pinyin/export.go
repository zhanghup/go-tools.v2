package pinyin

import (
	"strings"
)

func ExportPinyin(str string) string {
	args := NewArgs()
	args.Separator = ""
	s := Pinyin(str, args)
	result := ""
	for _, o := range s {
		for _, oo := range o {
			result += oo
		}
	}
	return result
}

func ExportPINYIN(str string) string {
	return strings.ToUpper(ExportPinyin(str))
}

func ExportPy(str string) string {
	args := NewArgs()
	args.Separator = ""
	s := Pinyin(str, args)
	result := ""
	for _, o := range s {
		for _, oo := range o {
			if len(oo) > 0 {
				result += oo[:1]
			}
		}
	}
	return result
}
func ExportPY(str string) string {
	return strings.ToUpper(ExportPy(str))
}
