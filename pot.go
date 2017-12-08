package pot

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
