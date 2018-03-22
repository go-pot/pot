package store

import (
	"fmt"
	"net/url"
)

var m = map[string]func(ur *url.URL) (Store, error){}

func RegisterStore(scheme string, f func(ur *url.URL) (Store, error)) {
	m[scheme] = f
}

func NewStore(sess string) (Store, error) {
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

	return store, nil
}
