package pot

import (
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Pot struct {
	negroni.Negroni                // 基础
	render          *render.Render // 渲染器
}

func NewPot(h ...negroni.Handler) *Pot {

	neg := negroni.New(
		negroni.NewRecovery(),
		negroni.NewLogger())

	if len(h) != 0 {
		neg = neg.With(h...)
	}
	p := &Pot{
		Negroni: *neg,
		render:  render.New(),
	}
	return p
}
