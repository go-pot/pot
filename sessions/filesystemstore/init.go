package filesystemstore

import (
	"net/url"

	"gopkg.in/pot.v1/sessions"
)

func init() {
	sessions.RegisterSession("file", func(ur *url.URL) (sessions.Store, error) {
		query := ur.Query()
		pairs := []byte(query.Get("keyPairs"))
		return New(query.Get("path"), pairs), nil
	})
}
