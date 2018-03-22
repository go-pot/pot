package redisstore

import (
	"net/url"
	"strconv"

	"gopkg.in/pot.v1/store"
)

func init() {
	store.RegisterStore("redis", func(ur *url.URL) (store.Store, error) {
		query := ur.Query()

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
		rs, err := NewRediStoreWithDB(int(size), network, address, password, db)
		rs.SetKeyPrefix(prefix)
		if err != nil {
			return nil, err
		}
		return rs, nil
	})
}
