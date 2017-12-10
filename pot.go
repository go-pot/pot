package pot // import "gopkg.in/pot.v1"

import (
	"gopkg.in/pot.v1/negroni"
)

type Pot struct {
	negroni.Negroni // 基础
}

func NewPot(h ...negroni.Handler) *Pot {
	neg := negroni.New(h...)

	p := &Pot{
		Negroni: *neg,
	}
	return p
}
