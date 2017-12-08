package pot

import (
	"net/http"

	"github.com/wzshiming/pot/sessions"
)

type Content struct {
	r        *http.Request
	w        http.ResponseWriter
	Session  sessions.Session
	Request  *Request
	Response *Response
}

func NewContent(w http.ResponseWriter, r *http.Request) *Content {
	return &Content{
		Response: NewResponse(w),
		Request:  NewRequest(r),
		Session:  sessions.GetSession(r),
		r:        r,
		w:        w,
	}
}

type PotHandler struct {
	Func func(cont *Content)
}

func NewPotHandler(f func(cont *Content)) *PotHandler {
	return &PotHandler{
		Func: f,
	}
}

func (h *PotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Func(NewContent(w, r))
}
