package redisstore

import (
	"testing"
	"time"

	ffmt "gopkg.in/ffmt.v1"
	"gopkg.in/pot.v1/store"
)

func TestA(t *testing.T) {
	s, err := store.NewStore("redis://db.gs.wzsm.studio:6379?db=3&password=&name=s&keyPairs=sjapp&keyPrefix=app:session_")
	if err != nil {
		ffmt.Mark(err)
		return
	}
	d, err := s.Load("3")
	ffmt.Mark(d)
	if err != nil {
		ffmt.Mark(err)
		return
	}

	s.Save("1", time.Now())

	s.Delete("2")
}
