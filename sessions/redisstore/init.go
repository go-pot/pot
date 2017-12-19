package redisstore

import (
	"net/url"
	"strconv"

	"gopkg.in/pot.v1/negroni"
	"gopkg.in/pot.v1/sessions"
)

func init() {
	sessions.RegisterSession("redis", func(ur *url.URL) (negroni.HandlerFunc, error) {
		query := ur.Query()
		pairs := []byte(query.Get("pairs"))
		if len(pairs) == 0 {
			pairs = []byte("pairs")
		}
		size, _ := strconv.ParseInt(query.Get("size"), 0, 0)
		if size <= 0 {
			size = 1024
		}
		network := query.Get("network")
		if network == "" {
			network = "tcp"
		}
		password := query.Get("password")
		address := ur.Host
		name := query.Get("name")
		if name == "" {
			name = "session"
		}
		rs, err := New(int(size), network, address, password, pairs)
		if err != nil {
			return nil, err
		}
		return sessions.Sessions(name, rs), nil
	})
}
