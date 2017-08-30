package session

import (
	"github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/redisstore"
)

func RedisStore(size int, network, address, password string, keyPairs ...[]byte) (sessions.Store, error) {
	return redisstore.New(size, network, address, password, keyPairs...)
}
