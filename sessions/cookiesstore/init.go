package cookiesstore

import (
	"net/url"

	"gopkg.in/pot.v1/negroni"
	"gopkg.in/pot.v1/sessions"
)

func init() {
	sessions.RegisterSession("cookie", func(ur *url.URL) (negroni.HandlerFunc, error) {
		query := ur.Query()
		pairs := []byte(query.Get("pairs"))
		name := query.Get("name")
		if name == "" {
			name = "session"
		}
		return sessions.Sessions(name, New(pairs)), nil
	})
}
