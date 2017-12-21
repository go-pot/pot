package sessions

import (
	"fmt"
	"net/url"

	"gopkg.in/pot.v1/negroni"
)

var m = map[string]func(ur *url.URL) (Store, error){}

func RegisterSession(scheme string, f func(ur *url.URL) (Store, error)) {
	m[scheme] = f
}

func NewSession(sess string, opt *Options) (negroni.HandlerFunc, error) {
	ur, err := url.Parse(sess)
	if err != nil {
		return nil, err
	}

	f, ok := m[ur.Scheme]
	if !ok {
		return nil, fmt.Errorf("Undefined %s", ur.Scheme)
	}

	store, err := f(ur)
	if err != nil {
		return nil, err
	}
	if opt != nil {
		store.Options(*opt)
	}
	name := ur.Query().Get("name")
	if name == "" {
		name = "session"
	}
	return Sessions(name, store), nil
}
