package pot

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func NewRequest(pot *Pot, r *http.Request) *Request {
	return &Request{
		pot:     pot,
		Request: *r,
	}
}

type Request struct {
	pot *Pot
	http.Request
	body []byte
}

func (c *Request) Origin() string {
	return c.Request.Header.Get("Origin")
}

func (c *Request) Host() string {
	return c.Request.Header.Get("Host")
}

func (c *Request) RealIP() string {
	ip := strings.Split(c.Request.Header.Get("X-Real-IP"), ":")[0]
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
	ip := strings.Split(c.Request.RemoteAddr, ":")
	if len(ip) > 0 {
		if ip[0] != "[" {
			return ip[0]
		}
	}
	return "127.0.0.1"
}

// RequestQueryForm 获取query内容 url编码
func (c *Request) QueryForm() url.Values {
	return c.Request.Form
}

// RequestBodyForm 获取body内容 url编码
func (c *Request) BodyForm() url.Values {
	return c.Request.PostForm
}

// RequestBody 获取body内容
func (c *Request) Body() (data []byte, err error) {
	if c.body != nil {
		return c.body, nil
	}
	defer func() {
		if data == nil {
			data = []byte{}
		}
		c.body = data
	}()

	r := c.Request.Body

	defer r.Close()
	return ioutil.ReadAll(r)
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
