package pot

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"net/url"

	"github.com/goincremental/negroni-sessions"
)

type Controller struct {
	Ctx  *Context
	Pot  *Pot
	body []byte
}

// GetSession 获取用户 session
func (c *Controller) GetSession() sessions.Session {
	return sessions.GetSession(c.Ctx.Request)
}

// RequestBody 获取body内容
func (c *Controller) RequestBody() (data []byte, err error) {
	if c.body != nil {
		return c.body, nil
	}
	defer func() {
		if data == nil {
			data = []byte{}
		}
		c.body = data
	}()

	r := c.Ctx.Request.Body

	defer r.Close()
	return ioutil.ReadAll(r)
}

// RequestCookie 获取cookie 内容
func (c *Controller) RequestCookie(name string) (string, error) {
	co, err := c.Ctx.Request.Cookie(name)
	if err != nil {
		return "", err
	}
	return co.Value, nil
}

// RequestQueryForm 获取query内容 url编码
func (c *Controller) RequestQueryForm() url.Values {
	return c.Ctx.Request.Form
}

// RequestBodyForm 获取body内容 url编码
func (c *Controller) RequestBodyForm() url.Values {
	return c.Ctx.Request.PostForm
}

// RequestBodyJSON 获取body内容 JSON编码
func (c *Controller) RequestBodyJSON(v interface{}) error {
	d, err := c.RequestBody()
	if err != nil {
		return err
	}
	return json.Unmarshal(d, v)
}

// RequestBodyXML 获取body内容 XML编码
func (c *Controller) RequestBodyXML(v interface{}) error {
	d, err := c.RequestBody()
	if err != nil {
		return err
	}
	return xml.Unmarshal(d, v)
}

// ResponseJSON 返回JSON
func (c *Controller) ResponseJSON(status int, v interface{}) error {
	return c.Pot.render.JSON(c.Ctx.ResponseWriter, status, v)
}

// ResponseJSONP 返回JSONP
func (c *Controller) ResponseJSONP(status int, callback string, v interface{}) error {
	return c.Pot.render.JSONP(c.Ctx.ResponseWriter, status, callback, v)
}

// ResponseXML 返回XML
func (c *Controller) ResponseXML(status int, v interface{}) error {
	return c.Pot.render.XML(c.Ctx.ResponseWriter, status, v)
}

// ResponseText 返回Text
func (c *Controller) ResponseText(status int, v string) error {
	return c.Pot.render.Text(c.Ctx.ResponseWriter, status, v)
}

// ResponseCsvFile 返回 csv 格式文件
func (b *Controller) ResponseCsvFile(status int, name string, ss [][]string) error {
	buf := bytes.NewBuffer(nil)
	err := csv.NewWriter(buf).WriteAll(ss)
	if err != nil {
		return err
	}
	b.Ctx.ResponseWriter.Header().Set("Content-type", "text/csv")
	b.Ctx.ResponseWriter.Header().Set("Content-disposition", "attachment;filename="+name+".csv")
	b.Ctx.ResponseWriter.WriteHeader(status)
	_, err = b.Ctx.ResponseWriter.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// Response 返回数据
func (c *Controller) Response(status int, v []byte) error {
	return c.Pot.render.Data(c.Ctx.ResponseWriter, status, v)
}
