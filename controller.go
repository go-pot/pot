package pot

import (
	"net/http"

	"github.com/wzshiming/pot/sessions"
)

type Controller struct {
	Pot      *Pot
	Session  sessions.Session
	Request  *Request
	Response *Response
}

func (c *Controller) Init(pot *Pot, w http.ResponseWriter, r *http.Request) {
	c.Pot = pot
	c.Session = sessions.GetSession(r)
	c.Request = NewRequest(pot, r)
	c.Response = NewResponse(pot, w)

}
