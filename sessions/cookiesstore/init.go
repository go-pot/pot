package cookiesstore

import (
	"net/url"

	"gopkg.in/pot.v1/sessions"
)

func init() {
	sessions.RegisterSession("cookie", func(ur *url.URL) (sessions.Store, error) {
		query := ur.Query()
		pairs := []byte(query.Get("keyPairs"))

		return New(pairs), nil
	})
}
