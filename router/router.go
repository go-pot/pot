package router

import (
	"fmt"
	"strings"
)

func PrintRouter(route *Route, router *Router, ancestors []*Route) error {
	p, _ := route.GetPathRegexp()
	qr, _ := route.GetQueriesRegexp()
	m, _ := route.GetMethods()
	met := "ANY"
	if len(m) != 0 {
		met = strings.Join(m, ",")
	}
	fmt.Println(met, p, strings.Join(qr, ","))
	return nil
}
