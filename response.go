package pot

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
)

const (
	// ContentBinary header value for binary data.
	ContentBinary = "application/octet-stream"
	// ContentHTML header value for HTML data.
	ContentHTML = "text/html"
	// ContentJSON header value for JSON data.
	ContentJSON = "application/json"
	// ContentJSONP header value for JSONP data.
	ContentJSONP = "application/javascript"
	// ContentLength header constant.
	ContentLength = "Content-Length"
	// ContentText header value for Text data.
	ContentText = "text/plain"
	// ContentType header constant.
	ContentType = "Content-Type"
	// ContentXML header value for XML data.
	ContentXML = "text/xml"
	// Default character encoding.
	defaultCharset = "; charset=UTF-8"
)

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		ResponseWriter: w,
		r:              r,
	}
}

type Response struct {
	http.ResponseWriter
	r *http.Request
}

// Redirect
func (c *Response) Redirect(status int, url string) {
	http.Redirect(c, c.r, url, status)
}

// ResponseJSON 返回JSON
func (c *Response) JSON(status int, v interface{}) error {
	c.Header().Set(ContentType, ContentJSON+defaultCharset)
	c.WriteHeader(status)
	return json.NewEncoder(c).Encode(v)
}

// ResponseJSONP 返回JSONP
func (c *Response) JSONP(status int, callback string, v interface{}) error {
	c.Header().Set(ContentType, ContentJSONP+defaultCharset)
	c.WriteHeader(status)
	_, err := c.Write([]byte(callback + "("))
	if err != nil {
		return err
	}
	err = json.NewEncoder(c).Encode(v)
	if err != nil {
		return err
	}
	_, err = c.Write([]byte(");"))
	if err != nil {
		return err
	}
	return nil
}

// ResponseXML 返回XML
func (c *Response) XML(status int, v interface{}) error {
	c.Header().Set(ContentType, ContentXML+defaultCharset)
	c.WriteHeader(status)
	return xml.NewEncoder(c).Encode(v)
}

// ResponseText 返回Text
func (c *Response) Text(status int, v string) error {
	c.Header().Set(ContentType, ContentText+defaultCharset)
	c.WriteHeader(status)
	_, err := c.Write([]byte(v))
	return err
}

// ResponseData 返回数据
func (c *Response) Data(status int, v []byte) error {
	c.Header().Set(ContentType, ContentBinary)
	c.WriteHeader(status)
	_, err := c.Write(v)
	return err
}
