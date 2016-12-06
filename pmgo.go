package pmgo

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Dialer interface {
	Dial(string) (SessionManager, error)
	DialWithInfo(*mgo.DialInfo) (SessionManager, error)
	DialWithTimeout(string, time.Duration) (SessionManager, error)
}

type dialer struct{}

func NewDialer() Dialer {
	return new(dialer)
}

func (d *dialer) Dial(url string) (SessionManager, error) {
	s, err := mgo.Dial(url)
	se := &Session{
		session: s,
	}
	return se, err
}

func (d *dialer) DialWithInfo(info *mgo.DialInfo) (SessionManager, error) {
	s, err := mgo.DialWithInfo(info)
	se := &Session{
		session: s,
	}
	return se, err
}

func (d *dialer) DialWithTimeout(url string, timeout time.Duration) (SessionManager, error) {
	s, err := mgo.DialWithTimeout(url, timeout)
	se := &Session{
		session: s,
	}
	return se, err
}
