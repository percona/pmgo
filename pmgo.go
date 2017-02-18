package pmgo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type Dialer interface {
	Dial(string) (SessionManager, error)
	DialWithInfo(*DialInfo) (SessionManager, error)
	DialWithTimeout(string, time.Duration) (SessionManager, error)
}

type DialInfo struct {
	CACertFile string
	/* ClientCertFile and ClientCertKey can be the same file, as long as it
	   has both, certificate and key in it, like:
	   -----BEGIN PRIVATE KEY-----
	   ABCDEF0123456789...
	   -----END PRIVATE KEY-----
	   -----BEGIN CERTIFICATE-----
	   9876543210FEDCBA...
	   -----END CERTIFICATE-----
	*/
	ClientCertFile string
	ClientKeyFile  string

	Addrs          []string
	Direct         bool
	Timeout        time.Duration
	FailFast       bool
	Database       string
	ReplicaSetName string
	Source         string
	Service        string
	ServiceHost    string
	Mechanism      string
	Username       string
	Password       string
	PoolLimit      int
	DialServer     func(addr *mgo.ServerAddr) (net.Conn, error)
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

func (d *dialer) DialWithInfo(info *DialInfo) (SessionManager, error) {
	mgoInfo := &mgo.DialInfo{
		Addrs:          info.Addrs,
		Direct:         info.Direct,
		Timeout:        info.Timeout,
		FailFast:       info.FailFast,
		Database:       info.Database,
		ReplicaSetName: info.ReplicaSetName,
		Source:         info.Source,
		Service:        info.Service,
		ServiceHost:    info.ServiceHost,
		Mechanism:      info.Mechanism,
		Username:       info.Username,
		Password:       info.Password,
		PoolLimit:      info.PoolLimit,
		DialServer:     info.DialServer,
	}

	if info.CACertFile != "" || info.ClientCertFile != "" || info.ClientKeyFile != "" {
		tlsConfig := &tls.Config{}

		if info.CACertFile != "" {
			if _, err := os.Stat(info.CACertFile); os.IsNotExist(err) {
				return nil, err
			}

			roots := x509.NewCertPool()
			var ca []byte
			var err error

			if ca, err = ioutil.ReadFile(info.CACertFile); err != nil {
				return nil, fmt.Errorf("invalid pem file: %s", err.Error())
			}
			roots.AppendCertsFromPEM(ca)
			tlsConfig.RootCAs = roots

		}

		if info.ClientCertFile != "" && info.ClientKeyFile != "" {
			cert, err := tls.LoadX509KeyPair(info.ClientCertFile, info.ClientKeyFile)
			if err != nil {
				return nil, err
			}
			tlsConfig.Certificates = []tls.Certificate{cert}
		}

		mgoInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			return conn, err
		}

		mgoInfo.Source = "$external"
		mgoInfo.Mechanism = "MONGODB-X509"
	}

	s, err := mgo.DialWithInfo(mgoInfo)

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
