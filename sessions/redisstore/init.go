package redisstore

import (
	"net/url"
	"strconv"

	"gopkg.in/pot.v1/sessions"
)

func init() {
	sessions.RegisterSession("redis", func(ur *url.URL) (sessions.Store, error) {
		query := ur.Query()
		pairs := []byte(query.Get("keyPairs"))
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
		db := query.Get("db")
		if db == "" {
			db = "0"
		}
		password := query.Get("password")
		address := ur.Host
		prefix := query.Get("keyPrefix")
		if prefix == "" {
			prefix = "session:"
		}
		rs, err := NewWithDB(int(size), network, address, password, db, prefix, pairs)
		if err != nil {
			return nil, err
		}
		return rs, nil
	})
}
