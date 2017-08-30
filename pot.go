package pot

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
)

type Pot struct {
	negroni.Negroni                // 基础
	render          *render.Render // 渲染器
}

func NewPot(h ...negroni.Handler) *Pot {
	p := &Pot{
		Negroni: *negroni.New(
			negroni.NewRecovery(),
			negroni.NewLogger()).With(h...),
		render: render.New(),
	}
	return p
}

// 启用 session
func (p *Pot) EnabledSession(name string, store sessions.Store) {
	p.With(sessions.Sessions(name, store))
}
