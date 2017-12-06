package pot

import (
	"github.com/unrolled/render"
)

type Pot struct {
	Negroni                // 基础
	render  *render.Render // 渲染器
}

func NewPot(h ...Handler) *Pot {
	neg := New(h...)

	p := &Pot{
		Negroni: *neg,
		render:  render.New(),
	}
	return p
}
