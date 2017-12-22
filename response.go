package pot

import (
	"net/http"

	"gopkg.in/pot.v1/render"
)

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		ResponseWriter: w,
		r:              r,
		render:         render.New(),
	}
}

type Response struct {
	http.ResponseWriter
	r      *http.Request
	render *render.Render
}

// Redirect
func (c *Response) Redirect(url string, code int) {
	http.Redirect(c, c.r, url, code)
}

// ResponseJSON 返回JSON
func (c *Response) JSON(status int, v interface{}) error {
	return c.render.JSON(c, status, v)
}

// ResponseJSONP 返回JSONP
func (c *Response) JSONP(status int, callback string, v interface{}) error {
	return c.render.JSONP(c, status, callback, v)
}

// ResponseXML 返回XML
func (c *Response) XML(status int, v interface{}) error {
	return c.render.XML(c, status, v)
}

// ResponseText 返回Text
func (c *Response) Text(status int, v string) error {
	return c.render.Text(c, status, v)
}

// ResponseData 返回数据
func (c *Response) Data(status int, v []byte) error {
	return c.render.Data(c, status, v)
}
