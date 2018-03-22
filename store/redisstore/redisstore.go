// Copyright 2012 Brian "bojo" Jones. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package redisstore

import (
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
	"gopkg.in/pot.v1/store"
)

// RediStore stores sessions in a redis backend.
type RediStore struct {
	Pool       *redis.Pool
	MaxAge     int //  MaxAge == 0 session
	maxLength  int
	keyPrefix  string
	serializer store.SessionSerializer
}

// SetMaxLength sets RediStore.maxLength if the `l` argument is greater or equal 0
// maxLength restricts the maximum length of new sessions to l.
// If l is 0 there is no limit to the size of a session, use with caution.
// The default for a new RediStore is 4096. Redis allows for max.
// value sizes of up to 512MB (http://redis.io/topics/data-types)
// Default: 4096,
func (s *RediStore) SetMaxLength(l int) {
	if l >= 0 {
		s.maxLength = l
	}
}

// SetKeyPrefix set the prefix
func (s *RediStore) SetKeyPrefix(p string) {
	s.keyPrefix = p
}

// SetSerializer sets the serializer
func (s *RediStore) SetSerializer(ss store.SessionSerializer) {
	s.serializer = ss
}

//// Default is the one provided by this package value - `sessionExpire`.
//// Set it to 0 for no restriction.
//// Because we use `MaxAge` also in SecureCookie crypting algorithm you should
//// use this function to change `MaxAge` value.
func (s *RediStore) SetMaxAge(v int) {
	s.MaxAge = v
}

func dial(network, address, password string) (redis.Conn, error) {
	c, err := redis.Dial(network, address)
	if err != nil {
		return nil, err
	}
	if password != "" {
		if _, err := c.Do("AUTH", password); err != nil {
			c.Close()
			return nil, err
		}
	}
	return c, err
}

// NewRediStore returns a new RediStore.
// size: maximum number of idle connections.
func NewRediStore(size int, network, address, password string) (*RediStore, error) {
	return NewRediStoreWithPool(&redis.Pool{
		MaxIdle:     size,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dial(network, address, password)
		},
	})
}

func dialWithDB(network, address, password, DB string) (redis.Conn, error) {
	c, err := dial(network, address, password)
	if err != nil {
		return nil, err
	}
	if _, err := c.Do("SELECT", DB); err != nil {
		c.Close()
		return nil, err
	}
	return c, err
}

// NewRediStoreWithDB - like NewRedisStore but accepts `DB` parameter to select
// redis DB instead of using the default one ("0")
func NewRediStoreWithDB(size int, network, address, password, DB string) (*RediStore, error) {
	return NewRediStoreWithPool(&redis.Pool{
		MaxIdle:     size,
		IdleTimeout: 240 * time.Second,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		Dial: func() (redis.Conn, error) {
			return dialWithDB(network, address, password, DB)
		},
	})
}

// NewRediStoreWithPool instantiates a RediStore with a *redis.Pool passed in.
func NewRediStoreWithPool(pool *redis.Pool) (*RediStore, error) {
	rs := &RediStore{
		Pool:       pool,
		MaxAge:     60 * 20, // 20 minutes seems like a reasonable default
		maxLength:  4096,
		keyPrefix:  "store_",
		serializer: store.JSONSerializer{},
	}
	_, err := rs.ping()
	return rs, err
}

// Close closes the underlying *redis.Pool
func (s *RediStore) Close() error {
	return s.Pool.Close()
}

// ping does an internal ping against a server to check if it is alive.
func (s *RediStore) ping() (bool, error) {
	conn := s.Pool.Get()
	defer conn.Close()
	data, err := conn.Do("PING")
	if err != nil || data == nil {
		return false, err
	}
	return (data == "PONG"), nil
}

// save stores the session in redis.
func (s *RediStore) Save(key string, d interface{}) error {

	b, err := s.serializer.Serialize(d)
	if err != nil {
		return err
	}
	if s.maxLength != 0 && len(b) > s.maxLength {
		return errors.New("SessionStore: the value to store is too big")
	}
	conn := s.Pool.Get()
	defer conn.Close()
	if err = conn.Err(); err != nil {
		return err
	}

	_, err = conn.Do("SETEX", s.keyPrefix+key, s.MaxAge, b)
	return err
}

// load reads the session from redis.
// returns true if there is a sessoin data in DB
func (s *RediStore) Load(key string, d interface{}) error {
	conn := s.Pool.Get()
	defer conn.Close()
	if err := conn.Err(); err != nil {
		return err
	}
	data, err := conn.Do("GET", s.keyPrefix+key)
	if err != nil {
		return err
	}
	if data == nil {
		return nil // no data was associated with this key
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return err
	}

	err = s.serializer.Deserialize(b, &d)
	if err != nil {
		return err
	}
	return nil
}

// delete removes keys from redis if MaxAge<0
func (s *RediStore) Delete(key string) error {

	conn := s.Pool.Get()
	defer conn.Close()
	if _, err := conn.Do("DEL", s.keyPrefix+key); err != nil {
		return err
	}
	return nil
}
