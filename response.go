package pot

import (
	"net/http"
)

func NewResponse(pot *Pot, w http.ResponseWriter) *Response {
	return &Response{
		pot:            pot,
		ResponseWriter: w,
	}
}

type Response struct {
	pot *Pot
	http.ResponseWriter
}

// ResponseJSON 返回JSON
func (c *Response) JSON(status int, v interface{}) error {
	return c.pot.render.JSON(c, status, v)
}

// ResponseJSONP 返回JSONP
func (c *Response) JSONP(status int, callback string, v interface{}) error {
	return c.pot.render.JSONP(c, status, callback, v)
}

// ResponseXML 返回XML
func (c *Response) XML(status int, v interface{}) error {
	return c.pot.render.XML(c, status, v)
}

// ResponseText 返回Text
func (c *Response) Text(status int, v string) error {
	return c.pot.render.Text(c, status, v)
}

// ResponseData 返回数据
func (c *Response) Data(status int, v []byte) error {
	return c.pot.render.Data(c, status, v)
}
