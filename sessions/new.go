package sessions

import (
	"fmt"
	"net/url"

	"gopkg.in/pot.v1/negroni"
)

var m = map[string]func(ur *url.URL) (negroni.HandlerFunc, error){}

func RegisterSession(scheme string, f func(ur *url.URL) (negroni.HandlerFunc, error)) {
	m[scheme] = f
}

func NewSession(sess string) (negroni.HandlerFunc, error) {
	ur, err := url.Parse(sess)
	if err != nil {
		return nil, err
	}

	f, ok := m[ur.Scheme]
	if !ok {
		return nil, fmt.Errorf("Undefined %s", ur.Scheme)
	}
	return f(ur)
}
