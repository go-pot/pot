package pot

import (
	"net/http"

	"gopkg.in/pot.v1/sessions"
)

type Content struct {
	r    *http.Request
	w    http.ResponseWriter
	sess sessions.Session
	req  *Request
	resp *Response
}

func NewContent(w http.ResponseWriter, r *http.Request) *Content {
	return &Content{
		r: r,
		w: w,
	}
}

func (s *Content) Request() *Request {
	if s.req == nil && s.r != nil && s.w != nil {
		s.req = NewRequest(s.w, s.r)
	}
	return s.req
}

func (s *Content) Response() *Response {
	if s.resp == nil && s.r != nil && s.w != nil {
		s.resp = NewResponse(s.w, s.r)
	}
	return s.resp
}

func (s *Content) Session() sessions.Session {
	if s.sess == nil {
		s.sess = sessions.GetSession(s.r)
	}
	return s.sess
}
