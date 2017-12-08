package pot

import (
	"net/http"
)

type Controller struct {
	Pot *Pot
	Content
}

func (c *Controller) Init(pot *Pot, w http.ResponseWriter, r *http.Request) {
	c.Pot = pot
	c.Content = *NewContent(w, r)

}
