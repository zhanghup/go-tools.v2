package tools

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zhanghup/go-tools.v2/lib/pinyin"
	"reflect"
	"regexp"
	"strconv"
	"sync"
)

var strfmtregex = regexp.MustCompile(`{{.*?}}`)
var fmtHanRegex = regexp.MustCompile("<<\\s*([\u4e00-\u9fa50-9a-zA-Z_]+)\\s*>>")

/*
	字符串格式化
	1. 允许 tools.Fmt(`hello << 世界 >>`,map[string]any{"世界":"world"}) => "hello world"
	2. 允许 tools.Fmt(`hello %s`,"world") => "hello world"
	3. 允许 tools.Fmt(`hello {{.name}}`,map[string]any{"name":"world"}) => "hello world"
*/
func Fmt(format string, args ...any) string {
	// 情况1先处理
	format = fmtHanRegex.ReplaceAllString(format, "{{ .$1 }}")

	pps := make([]any, 0)
	ppm := map[string]any{}

	for _, o := range args {
		vl := Rft.realValue(reflect.ValueOf(o))

		if vl.Kind() == reflect.Map {
			for _, oo := range vl.MapKeys() {
				if oo.Kind() == reflect.String {
					ppm[oo.String()] = vl.MapIndex(oo).Interface()
				}
			}
		} else {
			pps = append(pps, o)
		}
	}

	// 情况2再处理
	if len(pps) > 0 {
		format = fmt.Sprintf(format, pps...)
	}

	// 情况3处理
	if len(ppm) > 0 {
		format = TextTemplate(format, ppm).String()
	}

	return format
}

// JSONString 以json格式输出struct对象,format判断时间将json格式化
func JSONString(obj any, format ...bool) string {
	if obj == nil {
		return ""
	}
	var datas []byte
	if len(format) > 0 && format[0] {

		r, err := json.MarshalIndent(obj, "", "\t")
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	} else {
		r, err := json.Marshal(obj)
		if err != nil {
			datas = []byte("数据格式化异常")
		} else {
			datas = r
		}
	}
	return string(datas)
}

/*
	DataToBytes 将数据转换为[]byte
	支持的类型：
		bool/*bool/[]bool
		int8/*int8/[]int8
		uint8/*uint8/[]uint8
		int16/*int16/[]int16
		uint16/*uint16/[]uint16
		int32/*int32/[]int32
		uint32/*uint32/[]uint32
		int64/*int64/[]int64
		uint64/*uint64/[]uint64
		float32/*float32/[]float32
		float64/*float64/[]float64
*/
func DataToBytes[T any](n T) []byte {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, n)
	return bytesBuffer.Bytes()
}

/*
	BytesToData 将数据转换为自定义类型
*/
func BytesToData[T any](b []byte) T {
	bytesBuffer := bytes.NewBuffer(b)
	var x T
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return x
}

// Merge 将多个map合并为一个新的map
func Merge[Value any](m2 ...map[string]Value) map[string]Value {
	result := map[string]Value{}

	if len(m2) > 0 {
		for _, mm := range m2 {
			for k, v := range mm {
				result[k] = v
			}
		}
	}
	return result
}

func ToInt[T ~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64](s string) T {
	n, _ := strconv.ParseInt(s, 10, 64)
	return T(n)
}

func ToFloat[T ~float32 | ~float64](s string) T {
	n, _ := strconv.ParseFloat(s, 64)
	return T(n)
}

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

func PtrCheck(i any) error {
	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return nil
	}
	return errors.New("数据类型异常，必须为指针类型")
}

func Ptr[T any](v T) *T {
	return &v
}

func PtrOfUUID() *string {
	return Ptr(UUID())
}

func WaitPage(total, size int, fn func(page int)) {
	cnt := 0

	if size < 1 {
		cnt = total
	} else if total%size == 0 {
		cnt = total / size
	} else {
		cnt = int(total/size) + 1
	}

	Wait(cnt, fn)
}

//WaitRoutineN 同时开启多个Routine执行任务
func WaitRoutineN(n, dataLen int, fn func(routineN int, index int)) {
	g := sync.WaitGroup{}
	g.Add(dataLen)
	qu := Queue[int]{}
	for i := 0; i < dataLen; i++ {
		qu.Push(i)
	}

	for i := 0; i < n; i++ {
		go func(routineN int) {
			for {
				if ls := qu.Pop(1); len(ls) > 0 {
					fn(routineN, ls[0])
					g.Done()
				} else {
					break
				}

			}
		}(i)
	}

	g.Wait()
}
func Wait(n int, fn func(nn int)) {
	g := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		g.Add(1)
		go func(nn int) {
			fn(nn)
			g.Done()
		}(i)
	}

	g.Wait()
}

func Reverse[T any](ls []T) []T {
	for i, j := 0, len(ls)-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}
	return ls
}

func AnyToAny[T any](v ...T) []any {
	r := make([]any, 0)
	for _, vv := range v {
		r = append(r, vv)
	}
	return r
}
