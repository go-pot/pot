package session

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
)

func CookieStore(keyPairs ...[]byte) (sessions.Store, error) {
	return cookiestore.New(keyPairs...), nil
}
