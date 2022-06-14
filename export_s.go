package tools

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

// Equal 比较两个字符串数组是否相同
// flag: 是否顺序也必须相同
func Equal[T comparable](source []T, dist []T, flag ...bool) bool {
	if len(source) != len(dist) {
		return false
	}

	if len(flag) > 0 && flag[0] {
		for i := range source {
			if source[i] != dist[i] {
				return false
			}
		}

		return true
	}

	smap := map[T]struct{}{}
	for _, s := range source {
		smap[s] = struct{}{}
	}

	for _, s := range dist {
		_, ok := smap[s]
		if !ok {
			return false
		}
	}

	return true
}

// UUID 就是一个36位的UUID
func UUID() string {
	id := uuid.NewV4()
	return id.String()
}

// UUID_ 将 UUID 的”横杠“替代为”下划线“
func UUID_() string {
	id := uuid.NewV4()
	return strings.ReplaceAll(id.String(), "-", "_")
}

// Contains 判断字符串是否包含在数组中
func Contains[T comparable](src []T, tag T) bool {
	for _, s := range src {
		if s == tag {
			return true
		}
	}
	return false
}

// StrOfRand 取固定长度的随机字符串
// flag 自否可包含特殊字符
func StrOfRand(l int, flag ...bool) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		b := 0
		if len(flag) == 0 || !flag[0] {
			switch r.Int() % 3 {
			case 0:
				b = r.Intn(10) + 48
			case 1:
				b = r.Intn(26) + 65
			case 2:
				b = r.Intn(26) + 97
			}
		} else {
			b = r.Intn(90) + 33
		}
		res[i] = byte(b)
	}
	return string(res)
}
