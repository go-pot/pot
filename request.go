package pot

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func NewRequest(w http.ResponseWriter, r *http.Request) *Request {
	return &Request{
		Request: *r,
		w:       w,
	}
}

type Request struct {
	http.Request
	w    http.ResponseWriter
	body []byte
}

func (c *Request) Origin() string {
	return c.Request.Header.Get("Origin")
}

func (c *Request) Host() string {
	return c.Request.Header.Get("Host")
}

func (c *Request) RealIP() string {
	ip := strings.Split(c.Header.Get("X-Real-IP"), ":")[0]
	if ip == "" {
		ip = c.IP()
	}
	return ip
}

func (c *Request) Proxy() []string {
	ips := c.Request.Header.Get("X-Forwarded-For")
	if ips != "" {
		return strings.Split(ips, ",")
	}
	return []string{}
}

func (c *Request) IP() string {
	ips := c.Proxy()
	if len(ips) > 0 && ips[0] != "" {
		rip := strings.Split(ips[0], ":")
		return rip[0]
	}
	ip := strings.Split(c.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}

// BodyFile
func (c *Request) BodyFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return c.Request.FormFile(key)
}

// QueryFormValue 获取query内容 url编码
func (c *Request) QueryFormValue(key string) string {
	return c.Request.FormValue(key)
}

// RequestQueryForm 获取query内容 url编码
func (c *Request) QueryForm() url.Values {
	return c.Request.Form
}

// RequestBodyForm 获取body内容 url编码
func (c *Request) BodyForm() url.Values {
	return c.Request.PostForm
}

// VarsValue 获取Url {} 内的变量
func (c *Request) VarsValue(key string) string {
	return c.Vars()[key]
}

// Vars 获取Url {} 内的变量
func (c *Request) Vars() map[string]string {
	return mux.Vars(&c.Request)
}

// RequestBody 获取body内容
func (c *Request) Body() ([]byte, error) {
	if c.body != nil {
		return c.body, nil
	}

	r := c.Request.Body

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	r.Close()

	c.body = data
	return c.body, nil
}

// RequestBodyJSON 获取body内容 JSON编码
func (c *Request) BodyJSON(v interface{}) error {
	d, err := c.Body()
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}

// RequestBodyXML 获取body内容 XML编码
func (c *Request) BodyXML(v interface{}) error {
	d, err := c.Body()
	if err != nil {
		return err
	}
	return xml.Unmarshal(d, v)
}
