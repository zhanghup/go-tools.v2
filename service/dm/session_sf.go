package dm

import (
	"context"
	"fmt"
	"github.com/zhanghup/go-tools.v2"
	"reflect"
	"regexp"
	"strings"
	"xorm.io/xorm"
)

var (
	regexpRemoveSqlContent = regexp.MustCompile(`\(.*\)`)
	regexpFindOrderBy      = regexp.MustCompile(`order\s+by\s+`)
	regexpFindOrderByAsc   = regexp.MustCompile(`[a-zA-Z0-9_]+`)
	regexpFindOrderByDesc  = regexp.MustCompile(`^-[a-zA-Z0-9_]+`)
	regexpFindLimit        = regexp.MustCompile(`limit`)
	regexpFindSqlQuery     = regexp.MustCompile(`[0-9a-zA-Z_]*\:?`)
	regexpFindSqlParam     = regexp.MustCompile(`[0-9a-zA-Z_]*\:[0-9a-zA-Z_]+`)
)

type sessionSF[T any] struct {
	context context.Context

	tableName string

	querySql  string
	queryArgs []any

	sqlParam map[string]any
	sqlArgs  []any

	limit *int
	skip  *int

	withs   []string
	orderby []string

	templateFunction map[string]any
}

func newSessionSf[T any](db *xorm.Engine, ctx context.Context) *sessionSF[T] {
	tab := tools.RftTypeInfo(new(T))

	return &sessionSF[T]{
		context:          ctx,
		tableName:        db.GetTableMapper().Obj2Table(tab.Name),
		templateFunction: tools.Merge(templateFunctions(db, "___with_"), templateFunctions(db, "___templates_"), templateFunctions(db, "___contexts_")),
	}
}

func (s *sessionSF[T]) SF(sqlstr string, querys ...any) {
	s.querySql = sqlstr
	s.queryArgs = querys
}

/*

	示例1：
		sql = "select * from user where a = ? and b = ?"
		querys = []interface{}{"a","b"}
	示例2：
		sql = "select * from user where a = :a and b = ?"
		querys = []interface{}{"b",map[string]interface{}{"a":"a"}}
	示例3：
		sql = "where a = ?"
		querys = []interface{}{"b"}
		bean = models.User

	>>> select user.* from user where a = ?

	@orderFlag: 是否加入排序内容，一般只有在查询的时候需要排序
	@selectArg: 是否需要拼接成完整的SQL
*/
func (s *sessionSF[T]) SQL(orderFlag, selectArg bool, columns ...string) string {
	sqlstr := strings.TrimSpace(s.querySql)
	sqlstr = strings.ReplaceAll(sqlstr, ":?", "?")

	// sql模板参数格式化
	query := map[string]any{}
	for i := range s.queryArgs {
		ty := reflect.TypeOf(s.queryArgs[i])
		if ty.Kind() == reflect.Map {
			vl := reflect.ValueOf(s.queryArgs[i])
			for _, key := range vl.MapKeys() {
				v := vl.MapIndex(key)
				query[key.String()] = v.Interface()
			}
		} else {
			uid := tools.UUID_()
			sqlstr = strings.Replace(sqlstr, "?", ":"+uid, 1)
			query[uid] = s.queryArgs[i]
		}
	}

	// sql模板格式化
	m1 := map[string]any{
		"withs": func(name string) string {
			s.withs = append(s.withs, name)
			return fmt.Sprintf("__sql_with_%s", name)
		},
		"ctx": func(name string) string {
			return fmt.Sprintf("{{ ctx_%s .ctx }}", name)
		},
	}

	// 通用text/template模板
	sqlstr = tools.TextTemplate(sqlstr, query).FuncMap(tools.Merge(m1, s.templateFunction)).String()

	// context 模板
	sqlstr = tools.TextTemplate(sqlstr, map[string]any{"ctx": s.context}).FuncMap(s.templateFunction).String()

	if strings.Index(sqlstr, "select") == 0 || strings.Index(sqlstr, "SELECT") == 0 {
		if len(columns) > 0 {
			sqlstr = fmt.Sprintf("select %s from (%s) _", strings.Join(columns, ","), sqlstr)
		}
	} else if selectArg {
		column := "*"
		if len(columns) > 0 {
			column = strings.Join(columns, ",")
		}

		switch {
		case strings.Index(sqlstr, "limit") == 0,
			strings.Index(sqlstr, "where") == 0,
			strings.Index(sqlstr, "group") == 0,
			strings.Index(sqlstr, "order") == 0,
			sqlstr == "":
			sqlstr = fmt.Sprintf("select %s from %s %s", column, s.tableName, sqlstr)
		default:
			sqlstr = fmt.Sprintf("select %s from %s where %s", column, s.tableName, sqlstr)
		}
	}

	// withs 模板
	if len(s.withs) > 0 {
		// 去重
		with_header := "\n with recursive "
		withs := []string{}
		wmap := map[string]bool{}
		for _, w := range s.withs {
			wmap[w] = true
		}
		for k := range wmap {
			kk := tools.TextTemplate(fmt.Sprintf("{{ ___with_%s .ctx }}", k), map[string]any{"ctx": s.context}).FuncMap(s.templateFunction).String()
			withs = append(withs, fmt.Sprintf("__sql_with_%s as (%s)", k, kk))
		}

		sqlstr = with_header + strings.Join(withs, ",") + " " + sqlstr
		sqlstr = tools.TextTemplate(sqlstr, map[string]any{"ctx": s.context}).FuncMap(s.templateFunction).String()
	}

	if orderFlag && len(s.orderby) > 0 {
		res := regexpRemoveSqlContent.ReplaceAllString(sqlstr, "")
		match := regexpFindOrderBy.MatchString(res)

		orderBy := make([]string, 0)
		for _, sss := range s.orderby {
			if regexpFindOrderByDesc.MatchString(sss) {
				ss := strings.Replace(sss, "-", "", 1)
				orderBy = append(orderBy, ss+" desc")
			} else if regexpFindOrderByAsc.MatchString(sss) {
				orderBy = append(orderBy, sss+" asc")
			} else {
				orderBy = append(orderBy, sss+" ")
			}
		}
		if match {
			sqlstr += "," + strings.Join(orderBy, ",")
		} else {
			sqlstr += " order by " + strings.Join(orderBy, ",")
		}
	}

	if s.limit != nil {
		res := regexpRemoveSqlContent.ReplaceAllString(sqlstr, "")
		if strings.Index(res, "limit") == -1 {
			if s.skip != nil {
				sqlstr = fmt.Sprintf("%s limit %d,%d", sqlstr, *s.skip, *s.limit)
			} else {
				sqlstr = fmt.Sprintf("%s limit %d", sqlstr, *s.limit)
			}
		}
	}

	s.sqlParam = query
	return s.args(sqlstr)
}

func (s *sessionSF[T]) Limit(n int) *sessionSF[T] {
	s.limit = &n
	return s
}

func (s *sessionSF[T]) Skip(n int) *sessionSF[T] {
	s.skip = &n
	return s
}

func (s *sessionSF[T]) args(sqlstring string) string {

	r := regexp.MustCompile(`[0-9a-zA-Z_]*\:[0-9a-zA-Z_\?]+`)
	ss := r.FindAllString(sqlstring, -1)

	result := make([]any, 0)

	for _, item := range ss {

		value, ok := s.sqlParam[strings.Split(item, ":")[1]]
		if !ok {
			continue
		}
		newSql, args := s.sf_args_item(sqlstring, item, reflect.ValueOf(value))
		sqlstring = newSql
		result = append(result, args...)

	}
	s.sqlArgs = result
	return sqlstring
}

func (s *sessionSF[T]) sf_args_item(sqlstring, kk string, value reflect.Value) (string, []any) {
	results := make([]any, 0)

	if value.Kind() == reflect.Ptr && value.Pointer() != 0 {
		return s.sf_args_item(sqlstring, kk, value.Elem())
	} else if value.Kind() == reflect.Ptr && value.Pointer() == 0 {
		sqlstring = strings.Replace(sqlstring, kk, "?", 1)
		results = append(results, value.Interface())
	} else {
		switch {
		case strings.HasPrefix(kk, "like:"):
			sqlstring = strings.Replace(sqlstring, kk, "like concat('%',?,'%')", 1)
			results = append(results, value.Interface())
		case strings.HasPrefix(kk, "between:"):
			sqlstring = strings.Replace(sqlstring, kk, "between ? and ?", 1)
			if value.Type().Kind() == reflect.Array || value.Type().Kind() == reflect.Slice {
				if value.Len() > 0 {
					results = append(results, value.Index(0).Interface())
					if value.Len() > 1 {
						results = append(results, value.Index(1).Interface())
					} else {
						results = append(results, nil)
					}
				} else {
					results = append(results, nil, nil)
				}
			} else {
				results = append(results, nil, nil)
			}
		case strings.HasPrefix(kk, "in:"):
			if value.Type().Kind() == reflect.Array || value.Type().Kind() == reflect.Slice {
				ps := []string{}
				args := []any{}
				for i := 0; i < value.Len(); i++ {
					v := value.Index(i)
					ps = append(ps, "?")
					args = append(args, v.Interface())
				}

				sqlstring = strings.Replace(sqlstring, kk, fmt.Sprintf("(%s)", strings.Join(ps, ",")), 1)
				results = append(results, args...)
			} else {
				sqlstring = strings.Replace(sqlstring, kk, "(?)", 1)
				results = append(results, nil)
			}
		default:
			if value.Type().Kind() == reflect.Array || value.Type().Kind() == reflect.Slice {
				ps := []string{}
				args := []any{}
				for i := 0; i < value.Len(); i++ {
					v := value.Index(i)
					ps = append(ps, "?")
					args = append(args, v.Interface())
				}

				sqlstring = strings.Replace(sqlstring, kk, fmt.Sprintf("(%s)", strings.Join(ps, ",")), 1)
				results = append(results, args...)
			} else {
				sqlstring = strings.Replace(sqlstring, kk, "?", 1)
				results = append(results, value.Interface())
			}

		}
	}

	return sqlstring, results
}
