package dm

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"github.com/zhanghup/go-tools.v2/service/loader"
	"reflect"
	"regexp"
	"strings"
	"xorm.io/xorm"
)

type LoaderResultItem[T any] struct {
	Info T      `xorm:"extends"`
	Nid  string `xorm:"_B51e761c0"`
}

var loaderSqlFormatRegexp = regexp.MustCompile(`^\w+$`)

func sqlFormat(sqlstr, field string) string {

	sqlstr = regexp.MustCompile(`^prefix_\S+\s+`).ReplaceAllString(sqlstr, "")

	otherQuery := ""
	ss := strings.Split(field, "|")
	field = ss[0]
	if len(ss) > 1 {
		otherQuery = " and " + ss[1]
	}

	if strings.Index(sqlstr, "select") == -1 && loaderSqlFormatRegexp.MatchString(sqlstr) {
		sqlstr = tools.TextTemplate(`
			select {{ .table }}.*,{{ .table }}.{{ .field }} _B51e761c0 from {{ .table }} where {{ .table }}.{{ .field }} in:keys {{ .other }}
		`, map[string]any{
			"table": sqlstr,
			"field": field,
			"other": otherQuery,
		}).String()
	} else {
		if otherQuery == "" {
			otherQuery = "1 = 1"
		}
		sqlstr = fmt.Sprintf(`select s.*,s.%s _B51e761c0 from (%s) s where %s`, field, sqlstr, otherQuery)
	}
	return sqlstr
}

// Slice 查找数据库对象,ctx可以为nil
func sliceLoader[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, order []string, param ...any) loader.IObject[[]Result] {

	info := tools.RftTypeInfo(make([]Result, 0))

	sid := ""
	var sess ISession[LoaderResultItem[Result]]
	if ctx == nil || ctx.Value(CONTEXT_SESSION) == nil {
		sess = Session[LoaderResultItem[Result]](db, ctx)
		sid = "sid"
	} else {
		sess = Context[LoaderResultItem[Result]](db, ctx)
		sid = sess.Id()
	}

	sqlstr := sqlFormat(beanNameOrSql, field)
	key := fmt.Sprintf("sid: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s,%v", sid, sqlstr, tools.JSONString(param), info.PkgPath, info.FullName, order)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))

	return loader.Load[[]Result](key, func(keys []string) (map[string][]Result, error) {
		res, err := sess.SF(sqlFormat(beanNameOrSql, field), append([]any{map[string]any{"keys": keys}}, param...)...).Order(order...).Find()

		result := map[string][]Result{}

		if err != nil {
			return result, err
		}

		for _, o := range res {
			result[o.Nid] = append(result[o.Nid], o.Info)
		}

		return result, err
	})
}

// Slice 查找数据库对象,ctx可以为nil
func Slice[Result any](db *xorm.Engine, ctx context.Context, beanKey, beanNameOrSql string, field string, order []string, param ...any) ([]Result, error) {
	l := sliceLoader[Result](db, ctx, beanNameOrSql, field, order, param...)
	res, ok, err := l.Load(beanKey)
	if err != nil || !ok {
		return nil, err
	}
	return res, nil
}

// Info 查找数据库对象,ctx可以为nil
func infoLoader[Result any](db *xorm.Engine, ctx context.Context, beanNameOrSql string, field string, param ...any) loader.IObject[Result] {
	info := tools.RftTypeInfo(make([]Result, 0))

	sid := ""
	var sess ISession[LoaderResultItem[Result]]
	if ctx == nil || ctx.Value(CONTEXT_SESSION) == nil {
		sess = Session[LoaderResultItem[Result]](db, ctx)
		sid = "sid"
	} else {
		sess = Context[LoaderResultItem[Result]](db, ctx)
		sid = sess.Id()
	}

	sqlstr := sqlFormat(beanNameOrSql, field)
	key := fmt.Sprintf("sid: %s, sql: %s, param: %s, bean.pkg: %s,bean.name: %s", sid, sqlstr, tools.JSONString(param), info.PkgPath, info.FullName)
	if info.Name == "" {
		key += ",bean.json: " + tools.JSONString(reflect.New(info.Type).Interface())
	}
	key = tools.MD5([]byte(key))

	return loader.Load[Result](key, func(keys []string) (map[string]Result, error) {
		res, err := sess.SF(sqlstr, append([]any{map[string]any{"keys": keys}}, param...)...).Find()

		result := map[string]Result{}

		if err != nil {
			return result, err
		}

		for _, o := range res {
			result[o.Nid] = o.Info
		}

		return result, err
	})
}

// Info 根据id查找数据库对象,ctx可以为nil
func Info[Result any](db *xorm.Engine, ctx context.Context, beanKey, beanNameOrSql string, field string, param ...any) (*Result, error) {
	l := infoLoader[Result](db, ctx, beanNameOrSql, field, param...)
	res, ok, err := l.Load(beanKey)
	if err != nil || !ok {
		return nil, err
	}
	return &res, nil
}
