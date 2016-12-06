package pmgo

import mgo "gopkg.in/mgo.v2"

type SessionManager interface {
	BuildInfo() (info mgo.BuildInfo, err error)
	// Clone() SessionManager
	Close()
	// Copy() *SessionManager
	DB(name string) DatabaseManager
	DatabaseNames() (names []string, err error)
	// EnsureSafe(safe *Safe)
	// FindRef(ref *DBRef) *Query
	// Fsync(async bool) error
	// FsyncLock() error
	// FsyncUnlock() error
	// LiveServers() (addrs []string)
	// Login(cred *Credential) error
	// LogoutAll()
	// Mode() Mode
	// New() *SessionManager
	// Ping() error
	// Refresh()
	// ResetIndexCache()
	Run(cmd interface{}, result interface{}) error
	// Safe() (safe *Safe)
	// SelectServers(tags ...bson.D)
	// SetBatch(n int)
	// SetBypassValidation(bypass bool)
	// SetCursorTimeout(d time.Duration)
	// SetMode(consistency Mode, refresh bool)
	// SetPoolLimit(limit int)
	// SetPrefetch(p float64)
	// SetSafe(safe *Safe)
	// SetSocketTimeout(d time.Duration)
	// SetSyncTimeout(d time.Duration)
}

type Session struct {
	session *mgo.Session
}

func (s *Session) BuildInfo() (info mgo.BuildInfo, err error) {
	return s.session.BuildInfo()
}

func (s *Session) Close() {
	s.session.Close()
}

func (s *Session) DB(name string) DatabaseManager {
	d := &Database{
		db: s.session.DB(name),
	}

	return d
}

func (s *Session) DatabaseNames() (names []string, err error) {
	return s.session.DatabaseNames()
}

func (s *Session) Run(cmd interface{}, result interface{}) error {
	return s.session.Run(cmd, result)
}
