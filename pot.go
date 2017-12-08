package pot // import "gopkg.in/pot.v1"

type Pot struct {
	Negroni // 基础
}

func NewPot(h ...Handler) *Pot {
	neg := NewNegroni(h...)

	p := &Pot{
		Negroni: *neg,
	}
	return p
}
