package pot

import "net/http"

// 请求体 返回体
type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
}
