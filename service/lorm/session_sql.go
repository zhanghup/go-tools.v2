package lorm

import (
	"fmt"
	"github.com/zhanghup/go-tools"
	"regexp"
	"strings"
)

func (s *session[T]) _sql(orderFlag bool) string {
	if !orderFlag || len(s.orderby) == 0 {
		return s.sql
	}
	res := regexp.MustCompile(`\(.*\)`).ReplaceAllString(s.sql, "")
	match := regexp.MustCompile(`order\s+by\s+`).MatchString(res)

	orderBy := make([]string, 0)
	for _, s := range s.orderby {
		if regexp.MustCompile(`^-[a-zA-Z0-9_]+`).MatchString(s) {
			ss := strings.Replace(s, "-", "", 1)
			orderBy = append(orderBy, ss+" desc")
		} else if regexp.MustCompile(`[a-zA-Z0-9_]+`).MatchString(s) {
			orderBy = append(orderBy, s+" asc")
		} else {
			orderBy = append(orderBy, s+" ")
		}
	}
	if match {
		return s.sql + "," + strings.Join(orderBy, ",")
	} else {
		return s.sql + " order by " + strings.Join(orderBy, ",")
	}
}

func (s *session[T]) _sql_with() string {
	sqlwith := ""
	if len(s.withs) > 0 {
		// 去重
		with_header := "\n with recursive "
		withs := []string{}
		wmap := map[string]bool{}
		for _, w := range s.withs {
			wmap[w] = true
		}
		for k := range wmap {
			kk := tools.TextTemplate(fmt.Sprintf("{{ tmp_%s .ctx }}", k), map[string]any{"ctx": s.Context()}).FuncMap(getWiths(s.engine.db)).String()
			withs = append(withs, fmt.Sprintf("__sql_with_%s as (%s)", k, kk))
		}

		sqlwith = with_header + strings.Join(withs, ",")
		sqlwith = tools.TextTemplate(sqlwith, map[string]any{"ctx": s.Context()}).FuncMap(getWiths(s.engine.db)).String()
	}
	return sqlwith
}
